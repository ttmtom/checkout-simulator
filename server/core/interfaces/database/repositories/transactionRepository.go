package repositories

import (
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"crypto-checkout-simulator/server/core/models"
)

type Payment interface {
	CreateNewPayment(orderId int64, gatewayId string, paymentUrl string, amount float64) (*models.Payment, error)
	UpdatePaymentByNewEvent(orderId int64, status models.PaymentStatus, data paymentgateway.EventData) (*models.Payment, *models.PaymentEvent, error)
}
