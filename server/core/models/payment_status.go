package models

type PaymentStatus string

const (
	PaymentStatusNew       PaymentStatus = "NEW"
	PaymentStatusCreated   PaymentStatus = "CREATED"
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)
