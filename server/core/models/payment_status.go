package models

type PaymentStatus string

const (
	StatusNew       PaymentStatus = "NEW"
	StatusPending   PaymentStatus = "PENDING"
	StatusCompleted PaymentStatus = "COMPLETED"
	StatusSigned    PaymentStatus = "Signed"
	StatusFailed    PaymentStatus = "FAILED"
)
