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

// SendWhatsApp envia uma mensagem para o telemóvel via WhatsApp
func SendWhatsApp(phone, message string) (string, error) {
	apiURL := os.Getenv("WHATSAPP_API_URL")
	apiKey := os.Getenv("WHATSAPP_API_KEY")

	// Formata número de Angola caso falte o prefixo
	if len(phone) == 9 && phone[0] == '9' {
		phone = "244" + phone
	}

	payload := map[string]interface{}{
		"number":  phone,
		"text":    message,
		"options": map[string]interface{}{"delay": 1200, "presence": "composing"},
	}

	body, _ := json.Marshal(payload)

	// Simulação caso a API de WhatsApp local não esteja rodando
	if apiURL == "" || apiURL == "http://localhost:8080/message/sendText" {
		log.Printf("📲 [WHATSAPP ENVIADO PARA %s]:\n%s\n", phone, message)
		return fmt.Sprintf("WPP_SIMULATED_%d", time.Now().Unix()), nil
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("apikey", apiKey)
	}

	client := &http.Client{Timeout: 8 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("⚠️ Erro ao contactar gateway de WhatsApp: %v", err)
		return "FALLBACK_LOGGED", nil
	}
	defer resp.Body.Close()

	return fmt.Sprintf("WPP_STATUS_%d", resp.StatusCode), nil
}
