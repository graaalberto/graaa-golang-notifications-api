package models

import "time"

// SendNotificationRequest é o payload para disparar mensagens
type SendNotificationRequest struct {
	RecipientPhone string                 `json:"recipientPhone" binding:"required"` // Ex: "+244923000000"
	Channel        string                 `json:"channel"`                           // "whatsapp", "sms" ou "both"
	Type           string                 `json:"type" binding:"required"`           // trip_assigned, alert_critical, breakdown_alert, payment_confirmed
	Message        string                 `json:"message"`
	Metadata       map[string]interface{} `json:"metadata"`
}

// NotificationResponse resultado do envio
type NotificationResponse struct {
	ID               string    `json:"id"`
	Channel          string    `json:"channel"`
	Recipient        string    `json:"recipient"`
	Status           string    `json:"status"`
	SentAt           time.Time `json:"sentAt"`
	ProviderResponse string    `json:"providerResponse,omitempty"`
}

// NotificationLog é a entidade de banco de dados para histórico
type NotificationLog struct {
	ID        string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	Recipient string    `gorm:"not null;type:varchar(50)" json:"recipient"`
	Channel   string    `gorm:"not null;type:varchar(20)" json:"channel"` // whatsapp | sms | email
	Type      string    `gorm:"not null;type:varchar(50)" json:"type"`    // trip_assigned, trip_completed, alert_critical, payment_receipt
	Message   string    `gorm:"not null" json:"message"`
	Status    string    `gorm:"default:'sent';type:varchar(20)" json:"status"` // sent, delivered, failed
	SentAt    time.Time `json:"sentAt"`
}
