package controllers

import (
	"net/http"

	"github.com/akhill4054/room-backend/schemas"
	"github.com/gin-gonic/gin"
)

func GetChat(c *gin.Context) {
	type Uri struct {
		ChatId int `uri:"chatId" binding:"required"`
	}

	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

}

func ListChats(c *gin.Context) {

}
