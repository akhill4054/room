package v1

import "github.com/gin-gonic/gin"

func RegisterRoomRoutes(g *gin.RouterGroup) {
	g.GET("/rooms", foo)
	g.GET("/rooms/:id", foo)
	g.PATCH("/rooms/:id", foo)
	g.POST("/rooms", foo)
	g.DELETE("/rooms/:id", foo)
}

func foo(c *gin.Context) {

}
