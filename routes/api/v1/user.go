package v1

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(g *gin.RouterGroup) {
	g.GET("/users/:id", controllers.GetUser)
}
