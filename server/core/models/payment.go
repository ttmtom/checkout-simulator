package models

import (
	"time"
)

type Payment struct {
	ID                int64         `json:"id" db:"id"`
	OrderID           int64         `json:"orderId" db:"order_id"`
	ServiceProviderID string        `json:"serviceProviderId" db:"service_provider_id"`
	PaymentUrl        string        `json:"paymentUrl" db:"payment_url"`
	Status            PaymentStatus `json:"status" db:"status"`
	Amount            float64       `json:"amount,omitempty" db:"amount"`
	CreatedAt         time.Time     `json:"createdAt" db:"created_at"`
	LastEventAt       *time.Time    `json:"lastEventAt" db:"last_event_at"`
}
