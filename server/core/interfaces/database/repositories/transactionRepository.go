package repositories

import "crypto-checkout-simulator/server/core/models"

type Payment interface {
	Insert(payment *models.Payment) error
}
