package impl_repositories

import (
	"crypto-checkout-simulator/server/core/models"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		DB: db,
	}
}

func (t PaymentRepositoryImpl) Insert(payment *models.Payment) error {
	//TODO implement me
	panic("implement me")
}
