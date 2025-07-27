package payment

import (
	mockcoinbase "crypto-checkout-simulator/server/adapter/payment-gateway/mock-coinbase"
	"crypto-checkout-simulator/server/core/interfaces/database"
	"github.com/go-playground/validator"
)

type Module struct {
	controller *Controller
	service    *Service
}

func NewPaymentModule(store database.Storage, validator *validator.Validate) *Module {
	mockCoinbase := mockcoinbase.NewMockCoinbase()
	service := NewService(store.GetPaymentRepository(), store.GetOrderRepository(), mockCoinbase)
	controller := NewController(validator, service, mockCoinbase)

	return &Module{
		controller,
		service,
	}
}

func (m *Module) GetController() *Controller {
	return m.controller
}

func (m *Module) GetService() *Service {
	return m.service
}
