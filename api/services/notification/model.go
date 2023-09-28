package notification

import (
	"time"

	"golang.org/x/time/rate"
)

type Notification struct {
	ID      int16
	Type    string
	UserId  string
	Message string
}
type EmailSender struct {
	// Map to store scheduled rate limiters for each user
	rateLimiters map[string]*rate.Sometimes
	mailTypes    map[string]*MailTypeLimiter
}

type MailTypeLimiter struct {
	// Struct to store rate limiters rules for each type
	maxCount int16
	interval time.Duration
}
