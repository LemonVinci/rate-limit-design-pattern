package notification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SendEmail_Test(t *testing.T) {

	service := NewEmailSender()

	response := service.SendEmail("user1", "news", 1)
	assert.NotNil(t, response)
	assert.Equal(t, response[0:3], "200")

	response = service.SendEmail("user1", "news", 2)
	assert.NotNil(t, response)
	assert.Equal(t, response[0:3], "200")

	response = service.SendEmail("user1", "news", 2)
	assert.NotNil(t, response)
	assert.Equal(t, response[0:3], "429")
}
