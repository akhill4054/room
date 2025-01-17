package controllers

import (
	"net/http"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/schemas"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var reqBody schemas.SignUpRequestSchema

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			schemas.ErrorResponse{Message: err.Error()},
		)
	}

	user, err := models.CreateUser(
		reqBody.Username,
		reqBody.Email,
		reqBody.Password,
		reqBody.IsAdmin,
	)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == models.ErrUserPasswordValidation {
			statusCode = http.StatusBadRequest
		}
		c.IndentedJSON(statusCode, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	response := schemas.UserScheama{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
	c.IndentedJSON(http.StatusCreated, response)
}

func Login(c *gin.Context) {
	var reqBody schemas.LoginRequestSchema

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	user, err := models.GetUserByUsername(reqBody.Username)
	if err == models.ErrUserNotFound {
		c.IndentedJSON(
			http.StatusBadRequest,
			schemas.ErrorResponse{Message: models.ErrInvalidUsernameOrPassword.Error()},
		)
		return
	}

	token, err := models.GetAuthenticationToken(
		user.Id, reqBody.Username, reqBody.Password,
	)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	response := schemas.LoginResponseSchema{Token: token}
	c.IndentedJSON(http.StatusOK, response)
}
