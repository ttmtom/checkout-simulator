package repositories

import "crypto-checkout-simulator/server/core/models"

type TransactionRepository interface {
	Insert(transaction *models.Transaction) error
}
