package routes

import (
	"github.com/akhill4054/room-backend/middlewares"
	"github.com/akhill4054/room-backend/routes/api"
	v1 "github.com/akhill4054/room-backend/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api.RegisterAuthRoutes(r)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JwtAuthMiddleware())

	v1.RegisterUserRoutes(apiv1)
	v1.RegisterProifleRoutes(apiv1)
	v1.RegisterRoomRoutes(apiv1)
	v1.RegisterChatRoutes(apiv1)
	v1.RegisterMessageRoutes(apiv1)

	return r
}
