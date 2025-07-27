package models

import (
	"encoding/json"
	"time"
)

type PaymentEvent struct {
	ID                int64            `json:"id" db:"id"`
	ProviderType      string           `json:"providerType" db:"provider_type"`
	PaymentID         int64            `json:"paymentID" db:"payment_id"`
	ServiceProviderID string           `json:"serviceProviderId" db:"service_provider_id"`
	Status            OrderStatus      `json:"status" db:"status"`
	EventPayload      *json.RawMessage `json:"eventPayload,omitempty" db:"event_payload"`
	CreatedAt         time.Time        `json:"createdAt" db:"created_at"`
}
