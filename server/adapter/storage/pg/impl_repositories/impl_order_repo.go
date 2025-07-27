package impl_repositories

import (
	"crypto-checkout-simulator/server/core/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (t OrderRepositoryImpl) CreateNewOrder(userEmail string, amount float64) (*models.Order, error) {
	order := &models.Order{
		User:   userEmail,
		Amount: amount,
		Status: models.OrderStatusPending,
	}

	res := t.db.Create(&order)

	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}
