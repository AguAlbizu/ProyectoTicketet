package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// EmailClient define el contrato para enviar emails desde los servicios.
type EmailClient interface {
	SendEmail(to, subject, body string) error
}

// HTTPEmailClient envía emails a través de una API HTTP externa.
// Configuración por variables de entorno: EMAIL_API_URL, EMAIL_API_KEY, EMAIL_FROM.
type HTTPEmailClient struct {
	apiURL string
	apiKey string
	from   string
	client *http.Client
}

func NewEmailClient() *HTTPEmailClient {
	return &HTTPEmailClient{
		apiURL: os.Getenv("EMAIL_API_URL"),
		apiKey: os.Getenv("EMAIL_API_KEY"),
		from:   os.Getenv("EMAIL_FROM"),
		client: &http.Client{},
	}
}

type emailPayload struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (c *HTTPEmailClient) SendEmail(to, subject, body string) error {
	payload := emailPayload{To: to, From: c.from, Subject: subject, Body: body}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al serializar email: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.apiURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error al construir request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("la API de email retornó status %d", resp.StatusCode)
	}
	return nil
}

// NoOpEmailClient descarta los emails silenciosamente. Útil para desarrollo local.
type NoOpEmailClient struct{}

func (n *NoOpEmailClient) SendEmail(to, subject, body string) error { return nil }
