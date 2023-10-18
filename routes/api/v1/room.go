package v1

import (
	"github.com/akhill4054/room-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(g *gin.RouterGroup) {
	g.GET("/rooms", controllers.GetRooms)
	g.POST("/rooms", controllers.CreateRoom)
	g.GET("/rooms/:roomId", controllers.GetRoom)
	g.PUT("/rooms/:roomId", controllers.UpdateRoom)
	g.DELETE("/rooms/:roomId", controllers.DeleteRoom)

	g.POST("/rooms/:roomId/members", controllers.CreateRoomMember)
	g.DELETE("/rooms/:roomId/members/:memberId", controllers.DeleteRoomMember)
}
