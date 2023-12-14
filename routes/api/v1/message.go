package v1

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMessageRoutes(g *gin.RouterGroup) {
	g.POST("/messages", controllers.CreateMessage)
}
