package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// SendSMS envia mensagem SMS padrão (Unitel/Movicel)
func SendSMS(phone, message string) (string, error) {
	apiURL := os.Getenv("SMS_API_URL")
	apiKey := os.Getenv("SMS_API_KEY")
	sender := os.Getenv("SMS_SENDER_NAME")
	if sender == "" {
		sender = "FrotaGo"
	}

	payload := map[string]interface{}{
		"to":      phone,
		"from":    sender,
		"message": message,
	}

	body, _ := json.Marshal(payload)

	// Simulação/Log caso não tenha credenciais reais de operadora no momento
	if apiKey == "" || apiKey == "sms_angola_api_key_sandbox" {
		log.Printf("📩 [SMS ENVIADO PARA %s via %s]:\n%s\n", phone, sender, message)
		return fmt.Sprintf("SMS_SIMULATED_%d", time.Now().Unix()), nil
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("⚠️ Erro ao contactar gateway SMS: %v", err)
		return "SMS_FAILED", err
	}
	defer resp.Body.Close()

	return fmt.Sprintf("SMS_STATUS_%d", resp.StatusCode), nil
}
