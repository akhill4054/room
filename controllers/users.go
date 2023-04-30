package controllers

import (
	"net/http"
	"strconv"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/schemas"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := schemas.ErrorResponse{Message: "Invalid user id"}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	user, err := models.GetUser(userId)

	if err != nil {
		var statusCode = http.StatusInternalServerError
		if err == models.ErrUserNotFound {
			statusCode = http.StatusNotFound
		}
		c.IndentedJSON(statusCode, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	response := schemas.UserScheama{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	c.IndentedJSON(http.StatusOK, response)
}
