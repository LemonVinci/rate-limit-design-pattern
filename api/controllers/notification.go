package controllers

import (
	e "errors"
	"lemonvinci/rate-limiter-design-pattern/api/services/notification"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BurstNotification(c *gin.Context) {
	var newNotification notification.Notification

	ValidateParams(&newNotification, c)
	service := notification.NewEmailSender()

	// Simulate sending emails to different users
	userEmails := []string{"user1@example.com", "user2@example.com", "user3@example.com"}

	for i := 1; i <= 10; i++ {
		// Send emails to different users
		for _, userEmail := range userEmails {
			service.SendEmail(userEmail, "update", i)
			service.SendEmail(userEmail, "news", i)
		}

		// Sleep for a moment to simulate time passing
		time.Sleep(500 * time.Millisecond)
	}

	c.IndentedJSON(http.StatusCreated, newNotification)
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
