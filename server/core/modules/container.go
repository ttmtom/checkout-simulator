package modules

import (
	"crypto-checkout-simulator/server/core/interfaces/database"
	"crypto-checkout-simulator/server/core/modules/payment"
	"github.com/go-playground/validator"
)

type Container struct {
	paymentModule *payment.Module
}

func InitModuleContainer(storage *database.Storage, validator *validator.Validate) *Container {
	tm := payment.NewPaymentModule(storage, validator)

	return &Container{
		tm,
	}
}

func (c *Container) GetPaymentModule() *payment.Module {
	return c.paymentModule
}
