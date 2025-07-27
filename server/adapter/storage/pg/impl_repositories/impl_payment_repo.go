package impl_repositories

import (
	"crypto-checkout-simulator/server/core/models"
	"gorm.io/gorm"
	"time"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (t PaymentRepositoryImpl) CreateNewPayment(orderId int64, gatewayId string, paymentUrl string, amount float64) (*models.Payment, error) {
	payment := &models.Payment{
		OrderID:           orderId,
		ServiceProviderID: gatewayId,
		PaymentUrl:        paymentUrl,
		Status:            models.PaymentStatusNew,
		Amount:            amount,
		CreatedAt:         time.Now(),
	}

	res := t.db.Create(&payment)

	if res.Error != nil {
		return nil, res.Error
	}

	return payment, nil
}
