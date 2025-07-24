package models

import (
	"time"
)

type Transaction struct {
	ID                int64             `json:"id" db:"id"`
	ProviderType      string            `json:"providerType" db:"provider_type"`
	ServiceProviderID string            `json:"serviceProviderId" db:"service_provider_id"`
	CustomerEmail     *string           `json:"customerEmail,omitempty" db:"customer_email"`
	Status            TransactionStatus `json:"status" db:"status"`
	Amount            *float64          `json:"amount,omitempty" db:"amount"`
	Currency          *string           `json:"currency,omitempty" db:"currency"`
	CreatedAt         time.Time         `json:"createdAt" db:"created_at"`
	LastEventAt       time.Time         `json:"lastEventAt" db:"last_event_at"`
}
