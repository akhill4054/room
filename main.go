package main

import (
	"net/http"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/pkg/util"
	"github.com/akhill4054/room-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	util.SetupLogging()
	models.SetupDB()

	router := routes.InitRouter()

	router.GET("/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run("0.0.0.0:8000")
}
