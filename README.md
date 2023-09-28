# rate-limit-design-pattern
Rate limit solution to prevent burst of mails sent to recipient/userId

# Sample notification types and ratelimits rules
    "marketing": {1, 60 * time.Second}, up to 1 mails in 60 seconds
	"status":    {3, 20 * time.Second}, up to 3 mails in 20 seconds
	"news":      {2, 10 * time.Second}, up to 2 mails in 10 seconds
	"":          {1, 60 * time.Second}, up to 1 mails in 60 seconds (this type is the default rule for empty)

# Routers (defined in server.go)
GET /ping (returns a pong, used to check if API is serving HTTP on localhost:8080)

POST /notification (Send a single email to a user)
body: {
    "id": "1",
    "type": "news",
    "userId": "email@example.com",
    "message": "Mail del tipo News enviado al usuario 1111"
}


POST /burst_notifications (Simulate a burst of mails to different preset users in order to test each rate-limit rule)
body is not required

(WIP) POST /notification_type_limiter (create a notification type with custom rate-limit values)
