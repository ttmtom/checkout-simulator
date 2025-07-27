package payment

import (
	mockcoinbase "crypto-checkout-simulator/server/adapter/payment-gateway/mock-coinbase"
	"crypto-checkout-simulator/server/core/interfaces/database/repositories"
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"strconv"
)

type Service struct {
	paymentRepo repositories.Payment
	orderRepo   repositories.Order
	gateway     paymentgateway.Gateway
}

func NewService(paymentRepo repositories.Payment, orderRepo repositories.Order) *Service {

	return &Service{paymentRepo, orderRepo, mockcoinbase.NewMockCoinbase()}
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
