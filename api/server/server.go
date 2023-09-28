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
	router.POST("/notifications", controllers.SendNotification)
	router.POST("/burst_notification", controllers.BurstNotification)
	router.Run(domain)
}
