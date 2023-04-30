package api

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
}
