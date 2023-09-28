package server

import (
	"lemonvinci/rate-limiter-design-pattern/api/controllers"

	"github.com/gin-gonic/gin"
)

const domain = "localhost:8080"

var (
	Router *gin.Engine
)

func StartApp() {

	router := gin.Default()
	router.GET("/ping", controllers.Ping)
	router.POST("/notification", controllers.SendNotification)
	router.POST("/burst_notifications", controllers.BurstNotification)
	router.Run(domain)
}
