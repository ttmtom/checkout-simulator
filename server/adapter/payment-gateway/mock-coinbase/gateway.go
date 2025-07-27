package mock_coinbase

import (
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"fmt"
)

type MockCoinbase struct{}

func NewMockCoinbase() paymentgateway.Gateway {
	return &MockCoinbase{}
}

func (m *MockCoinbase) CreateCharge(id string) *paymentgateway.ChargeResponse {
	return &paymentgateway.ChargeResponse{
		Id:         id,
		PaymentUrl: fmt.Sprintf("https://fake.coinbase.com/pay/%s", id),
	}
}
