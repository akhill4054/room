package v1

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProifleRoutes(g *gin.RouterGroup) {
	g.POST("/profiles", controllers.CreateProfile)
	g.GET("/profiles", controllers.GetUserProfile)
	g.PUT("/profiles/profileId", controllers.UpdateUserProfile)
}
