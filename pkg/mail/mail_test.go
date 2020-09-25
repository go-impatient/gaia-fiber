package mail

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServerName tests with a port greater than 0.
func TestServerName(t *testing.T) {
	port := 2000
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost:2000"

	assert.Equal(t, got, wanted)
}

// TestServerNameWithoutPort tests with a port equal to 0.
func TestServerNameWithoutPort(t *testing.T) {
	port := 0
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost"

	assert.Equal(t, got, wanted)
}

// TestBuildMessageWithoutFrom tests BuildMessage without mail from.
func TestBuildMessageWithoutFrom(t *testing.T) {
	mail := Mail{
		To: make([]string, 1),
	}

	_, err := mail.buildMessage()

	assert.NotNil(t, err)
}

// TestBuildMessageWithouTo tests BuildMessage without mail to.
func TestBuildMessageWithoutTo(t *testing.T) {
	mail := Mail{
		From: "bob@alice.test",
	}

	_, err := mail.buildMessage()

	assert.NotNil(t, err)
}

// TestBuildMessage tests BuildMessage with good parameters.
func TestBuildMessage(t *testing.T) {
	to := make([]string, 2)
	to[0] = "alice@bob.test"
	to[1] = "john@doe.test"
	mail := Mail{
		From:    "bob@alice.test",
		To:      to,
		Subject: "Test",
		Body:    "Body",
	}

	got, _ := mail.buildMessage()
	wanted := fmt.Sprint("From: bob@alice.test\r\n")
	wanted += fmt.Sprint("To: alice@bob.test,john@doe.test\r\n")
	wanted += fmt.Sprint("Subject: Test\r\n")
	wanted += fmt.Sprint("MIME-version: 1.0\r\n")
	wanted += fmt.Sprint("Content-Type: text/html; charset: UTF-8\r\n")
	wanted += fmt.Sprint("\r\nBody")

	assert.Equal(t, got, wanted)
}

// TestSend tests mail send construction.
func TestSend(t *testing.T) {
	err := Send("from", []string{"test@example.com"}, "subject", "body", "smptUser", "smtpPassword", "smtpHost", 1234)

	assert.NotNil(t, err)
}

// TestSendMail tests mail send.
func TestSendMail(t *testing.T) {
	err := sendMail(Mail{}, "message", "", 1234)

	assert.NotNil(t, err)
}
