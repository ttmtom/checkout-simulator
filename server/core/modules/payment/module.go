package payment

import (
	"crypto-checkout-simulator/server/core/interfaces/database"
	"github.com/go-playground/validator"
)

type Module struct {
	controller *Controller
	service    *Service
}

func NewPaymentModule(store database.Storage, validator *validator.Validate) *Module {
	service := NewService(store.GetPaymentRepository(), store.GetOrderRepository())
	controller := NewController(validator, service)

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
