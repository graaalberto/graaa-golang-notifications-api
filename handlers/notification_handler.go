package handlers

import (
	"fmt"
	"net/http"
	"time"

	"frotago-notifications-api/config"
	"frotago-notifications-api/models"
	"frotago-notifications-api/services"

	"github.com/gin-gonic/gin"
)

// SendNotification dispara a mensagem por WhatsApp, SMS ou ambos
func SendNotification(c *gin.Context) {
	var req models.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Dados inválidos: " + err.Error()})
		return
	}

	if req.Message == "" {
		req.Message = BuildMessageFromTemplate(req.Type, req.Metadata)
	}

	channel := req.Channel
	if channel == "" {
		channel = "whatsapp"
	}

	notifID := fmt.Sprintf("ntf_%d", time.Now().UnixNano()%1000000)
	var providerResp string

	if channel == "whatsapp" || channel == "both" {
		resp, _ := services.SendWhatsApp(req.RecipientPhone, req.Message)
		providerResp += "[WhatsApp: " + resp + "]"
	}

	if channel == "sms" || channel == "both" {
		resp, _ := services.SendSMS(req.RecipientPhone, req.Message)
		providerResp += "[SMS: " + resp + "]"
	}

	// Grava no log de histórico do PostgreSQL
	logEntry := models.NotificationLog{
		ID:        notifID,
		Recipient: req.RecipientPhone,
		Channel:   channel,
		Type:      req.Type,
		Message:   req.Message,
		Status:    "sent",
		SentAt:    time.Now(),
	}
	config.DB.Create(&logEntry)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Notificação enviada com sucesso",
		"data": models.NotificationResponse{
			ID:               notifID,
			Channel:          channel,
			Recipient:        req.RecipientPhone,
			Status:           "sent",
			SentAt:           logEntry.SentAt,
			ProviderResponse: providerResp,
		},
	})
}

// GetNotificationLogs lista histórico de mensagens enviadas
func GetNotificationLogs(c *gin.Context) {
	var logs []models.NotificationLog
	if err := config.DB.Order("sent_at desc").Limit(100).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Erro ao buscar logs"})
		return
	}
	c.JSON(http.StatusOK, logs)
}
