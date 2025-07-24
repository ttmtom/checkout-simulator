package models

import (
	"encoding/json"
	"time"
)

type TransactionEvent struct {
	ID                int64             `json:"id" db:"id"`
	ProviderType      string            `json:"providerType" db:"provider_type"`
	TransactionID     int64             `json:"transactionId" db:"transaction_id"`
	ServiceProviderID string            `json:"serviceProviderId" db:"service_provider_id"`
	Status            TransactionStatus `json:"status" db:"status"`
	EventPayload      *json.RawMessage  `json:"eventPayload,omitempty" db:"event_payload"`
	CreatedAt         time.Time         `json:"createdAt" db:"created_at"`
}
