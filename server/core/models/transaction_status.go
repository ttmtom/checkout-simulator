package models

type TransactionStatus string

const (
	StatusNew        TransactionStatus = "NEW"
	StatusPending    TransactionStatus = "PENDING"
	StatusCompleted  TransactionStatus = "COMPLETED"
	StatusExpired    TransactionStatus = "EXPIRED"
	StatusUnresolved TransactionStatus = "UNRESOLVED"
	StatusCanceled   TransactionStatus = "CANCELED"
	StatusFailed     TransactionStatus = "FAILED"
)
