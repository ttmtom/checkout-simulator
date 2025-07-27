package payment

import (
	"crypto-checkout-simulator/server/core/interfaces/database/repositories"
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"crypto-checkout-simulator/server/core/models"
	"strconv"
)

type Service struct {
	paymentRepo repositories.Payment
	orderRepo   repositories.Order
	gateway     paymentgateway.Gateway
}

func NewService(paymentRepo repositories.Payment, orderRepo repositories.Order, gateway paymentgateway.Gateway) *Service {

	return &Service{paymentRepo, orderRepo, gateway}
}

type CheckoutResponse struct {
	PaymentUrl string `json:"payment_url"`
	ID         int64  `json:"id"`
}

func (s *Service) Checkout(email string, amount float64) (*CheckoutResponse, error) {
	order, err := s.orderRepo.CreateNewOrder(email, amount)
	if err != nil {
		return nil, err
	}

	coinbaseCharge := s.gateway.CreateCharge(strconv.FormatInt(order.ID, 10))
	_, err = s.paymentRepo.CreateNewPayment(order.ID, coinbaseCharge.Id, coinbaseCharge.PaymentUrl, amount)
	if err != nil {
		return nil, err
	}

	return &CheckoutResponse{
		PaymentUrl: coinbaseCharge.PaymentUrl,
		ID:         order.ID,
	}, nil
}

func (s *Service) WebhookPaymentCreatedHandler(data paymentgateway.EventData) {
	// when system receive created event, update the order status to processing, update  the payment status to created
	order, err := s.orderRepo.UpdateOrderStatus(data.Data.Metadata["order_id"].(int64), models.OrderStatusProcessing)
	if err != nil {
		return
	}

	s.paymentRepo.UpdatePaymentByNewEvent(order.ID, models.PaymentStatusCreated, data)
}

func (s *Service) WebhookPaymentConfirmedHandler(data paymentgateway.EventData) {
	order, err := s.orderRepo.UpdateOrderStatus(data.Data.Metadata["order_id"].(int64), models.OrderStatusCompleted)
	if err != nil {
		return
	}
	s.paymentRepo.UpdatePaymentByNewEvent(order.ID, models.PaymentStatusCompleted, data)

}

func (s *Service) WebhookPaymentFailedHandler(data paymentgateway.EventData) {
	order, err := s.orderRepo.UpdateOrderStatus(data.Data.Metadata["order_id"].(int64), models.OrderStatusFailed)
	if err != nil {
		return
	}
	s.paymentRepo.UpdatePaymentByNewEvent(order.ID, models.PaymentStatusFailed, data)
}
