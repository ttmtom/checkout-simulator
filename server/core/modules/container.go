package modules

import (
	"crypto-checkout-simulator/config"
	"crypto-checkout-simulator/server/adapter/storage/pg"
	"crypto-checkout-simulator/server/core/modules/payment"
	"github.com/go-playground/validator"
)

type Container struct {
	paymentModule *payment.Module
}

func InitModuleContainer(c *config.Config, validator *validator.Validate) *Container {
	storage, err := pg.New(c.Database)
	if err != nil {
		panic(err)
	}
	tm := payment.NewPaymentModule(storage, validator)

	return &Container{
		tm,
	}
}

func (c *Container) GetPaymentModule() *payment.Module {
	return c.paymentModule
}
