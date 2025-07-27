package repositories

import "crypto-checkout-simulator/server/core/models"

type Payment interface {
	CreateNewPayment(orderId int64, gatewayId string, paymentUrl string, amount float64) (*models.Payment, error)
}
