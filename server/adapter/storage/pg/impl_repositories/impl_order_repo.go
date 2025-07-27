package impl_repositories

import (
	"crypto-checkout-simulator/server/core/interfaces/database/repositories"
	"crypto-checkout-simulator/server/core/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) repositories.Order {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (o OrderRepositoryImpl) CreateNewOrder(userEmail string, amount float64) (*models.Order, error) {
	order := &models.Order{
		User:   userEmail,
		Amount: amount,
		Status: models.OrderStatusPending,
	}

	res := o.db.Create(&order)

	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}

func (o OrderRepositoryImpl) UpdateOrderStatus(orderId int64, status models.OrderStatus) (*models.Order, error) {
	res := o.db.Model(&models.Order{}).Where("id = ?", orderId).Update("status", status)
	if res.Error != nil {
		return nil, res.Error
	}

	var order models.Order
	o.db.Model(&models.Order{}).
		Select("*").
		Where("id = ?", res.RowsAffected).
		First(&order)

	return &order, nil
}
