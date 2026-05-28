package clients

import (
	"net/smtp"
)

// EmailClient wraps SMTP configuration for sending notification emails.
type EmailClient struct {
	Host     string
	Port     string
	Username string
	Password string
}

// NewEmailClient builds an EmailClient from environment variables.
// Reads SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD.
func NewEmailClient() *EmailClient {
	// TODO: read SMTP config from env and return populated struct
	_ = smtp.PlainAuth
	return &EmailClient{}
}

// SendRaffleWinnerEmail sends a congratulations email to the raffle winner.
func (c *EmailClient) SendRaffleWinnerEmail(to, eventName, raffleName string) error {
	// TODO: compose winner email body with eventName and raffleName
	// TODO: dial SMTP server and send using smtp.SendMail
	return nil
}

// SendRaffleLoserEmail sends a participation thank-you email to non-winners.
func (c *EmailClient) SendRaffleLoserEmail(to, eventName string) error {
	// TODO: compose consolation email body with eventName
	// TODO: dial SMTP server and send using smtp.SendMail
	return nil
}
