package controllers

import (
	e "errors"
	"lemonvinci/rate-limiter-design-pattern/api/services/notification"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BurstNotification(c *gin.Context) {
	const mailsToSend = 12

	service := notification.NewEmailSender()

	userEmails := []string{"user1@example.com", "user2@example.com"}
	logResponse := make([]string, 0)

	for i := 1; i <= mailsToSend; i++ {
		// Send emails to different users
		for _, userEmail := range userEmails {
			logResponse = append(logResponse, service.SendEmail(userEmail, "status", i))
			logResponse = append(logResponse, service.SendEmail(userEmail, "news", i))
			logResponse = append(logResponse, service.SendEmail(userEmail, "marketing", i))
		}

		// Sleep for a second to simulate time passing
		time.Sleep(time.Second)
	}

	c.IndentedJSON(http.StatusCreated, logResponse)
}

func SendNotification(c *gin.Context) {
	var newNotification notification.Notification

	ValidateParams(&newNotification, c)

	service := notification.NewEmailSender()

	service.SendEmail(newNotification.UserId, newNotification.Type, int(newNotification.ID))

	c.IndentedJSON(http.StatusCreated, newNotification)
}

func ValidateParams(newNotification *notification.Notification, c *gin.Context) {
	if err := c.BindJSON(&newNotification); err != nil {
		return
	}

	if newNotification.UserId == "" {
		_ = c.Error(e.New("userId cannot be empty"))
	}

	if newNotification.Type == "" {
		_ = c.Error(e.New("type cannot be empty"))
	}

	if err := c.BindJSON(&newNotification); err != nil {
		return
	}
}
