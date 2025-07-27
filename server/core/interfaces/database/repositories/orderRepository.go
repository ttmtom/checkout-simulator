package repositories

import "crypto-checkout-simulator/server/core/models"

type Order interface {
	CreateNewOrder(userEmail string, amount float64) (*models.Order, error)
}
