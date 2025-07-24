package impl_repositories

import (
	"crypto-checkout-simulator/server/core/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepoImpl(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

func (t TransactionRepositoryImpl) Insert(transaction *models.Transaction) error {
	//TODO implement me
	panic("implement me")
}
