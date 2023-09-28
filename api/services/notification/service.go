package notification

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

var mailTypes = map[string]MailType{
	"news":   MailType{1, 1000 * time.Millisecond},
	"update": MailType{2, 2000 * time.Millisecond},
	"":       MailType{3, 2000 * time.Millisecond},
}

type EmailSender struct {
	// Map to store scheduled rate limiters for each user
	rateLimiters map[string]*rate.Sometimes
	mailTypes    map[string]*MailType
}

type MailType struct {
	// Struct to store rate limiters rules for each type
	maxCount int16
	interval time.Duration
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		rateLimiters: make(map[string]*rate.Sometimes),
		mailTypes:    make(map[string]*MailType),
	}
}

func (es *EmailSender) SendEmail(userId string, mailType string, notificationId int) {
	var hasBeenSended bool

	// Get or create a scheduled rate limiter for the user
	limiter := es.getRateLimiter(userId, mailType)

	// Try to send an email
	limiter.Do(func() {
		fmt.Printf("%s email %d sent to %s\n", mailType, notificationId, userId)
		hasBeenSended = true
	})

	if !hasBeenSended {
		fmt.Printf("429 - Rate limit exceeded for %s. Waiting to send more %s emails...\n", mailType, userId)
	}

}

func (es *EmailSender) getRateLimiter(userId string, mailType string) *rate.Sometimes {
	// Check if a scheduled rate limiter exists for the hash userMail + mailType
	if limiter, ok := es.rateLimiters[userId+mailType]; ok {
		return limiter
	}

	limiter := rate.Sometimes{First: int(mailTypes[mailType].maxCount), Interval: mailTypes[mailType].interval}

	// Store the rate limiter in the map for the user
	es.rateLimiters[userId+mailType] = &limiter

	return &limiter
}
