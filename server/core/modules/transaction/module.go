package transaction

import (
	"crypto-checkout-simulator/server/core/interfaces/database"
	"github.com/go-playground/validator"
)

type Module struct {
	controller *Controller
	service    *Service
}

func NewTransactionModule(store *database.Storage, validator *validator.Validate) *Module {

	return &Module{
		controller: NewController(),
		service:    NewService(),
	}
}

func (m *Module) GetController() *Controller {
	return m.controller
}

func (m *Module) GetService() *Service {
	return m.service
}
