package v1

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(g *gin.RouterGroup) {
	g.POST("/chats", controllers.CreateRoom)
}
