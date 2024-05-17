package health

import (
	"fmt"
	"net/smtp"
	"os"
)

// EmailNotifier is responsible for sending email notifications.
type EmailNotifier struct {
	smtpServer string
	smtpPort   string
	username   string
	password   string
	fromEmail  string
}

// NewEmailNotifier creates a new EmailNotifier.
func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{
		smtpServer: os.Getenv("SMTP_SERVER"),
		smtpPort:   os.Getenv("SMTP_PORT"),
		username:   os.Getenv("SMTP_USERNAME"),
		password:   os.Getenv("SMTP_PASSWORD"),
		fromEmail:  os.Getenv("FROM_EMAIL"),
	}
}

// SendNotification sends an email notification to the specified recipient.
func (e *EmailNotifier) SendNotification(toEmail, subject, body string) error {
	auth := smtp.PlainAuth("", e.username, e.password, e.smtpServer)
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", e.fromEmail, toEmail, subject, body)
	addr := fmt.Sprintf("%s:%s", e.smtpServer, e.smtpPort)
	return smtp.SendMail(addr, auth, e.fromEmail, []string{toEmail}, []byte(msg))
}
