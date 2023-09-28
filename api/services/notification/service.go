package notification

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

var mailTypes = map[string]MailTypeLimiter{
	"marketing": {1, 60 * time.Second},
	"status":    {3, 20 * time.Second},
	"news":      {2, 10 * time.Second},
	"":          {1, 60 * time.Second},
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		rateLimiters: make(map[string]*rate.Sometimes),
		mailTypes:    make(map[string]*MailTypeLimiter),
	}
}

func (es *EmailSender) SendEmail(userId string, mailType string, notificationId int) string {
	var sentMessage string

	limiter := es.getSingletonRateLimiter(userId, mailType)

	limiter.Do(func() {
		fmt.Printf("200 - %s email %d sent to %s\n", mailType, notificationId, userId)
		sentMessage = fmt.Sprintf("200 - %s email %d sent to %s", mailType, notificationId, userId)
	})

	if sentMessage == "" {
		fmt.Printf("429 - %s email %d exceeded ratelimit for user %s.\n", mailType, notificationId, userId)
		return fmt.Sprintf("429 - %s email %d exceeded ratelimit for user %s.", mailType, notificationId, userId)
	}

	return sentMessage
}

func (es *EmailSender) getSingletonRateLimiter(userId string, mailType string) *rate.Sometimes {
	// Create a hash with userMail + mailType
	hashUserIdMailType := strings.ToLower(userId) + strings.ToLower(mailType)

	// Check if a scheduled rate limiter exists for the hash
	if limiter, ok := es.rateLimiters[hashUserIdMailType]; ok {
		return limiter
	}

	newLimiter := rate.Sometimes{First: int(mailTypes[mailType].maxCount), Interval: mailTypes[mailType].interval}

	// Store the rate limiter in the map for the user+type
	es.rateLimiters[hashUserIdMailType] = &newLimiter

	return &newLimiter
}
