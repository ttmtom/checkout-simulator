package modules

import (
	"crypto-checkout-simulator/server/core/interfaces/database"
	"crypto-checkout-simulator/server/core/modules/transaction"
	"github.com/go-playground/validator"
)

type Container struct {
	transactionModule *transaction.Module
}

func InitModuleContainer(storage *database.Storage, validator *validator.Validate) *Container {
	tm := transaction.NewTransactionModule(storage, validator)

	return &Container{
		tm,
	}
}

func (c *Container) GetTransactionModule() *transaction.Module {
	return c.transactionModule
}
