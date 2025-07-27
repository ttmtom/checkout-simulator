package impl_repositories

import (
	"context"
	"crypto-checkout-simulator/server/core/interfaces/database/repositories"
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"crypto-checkout-simulator/server/core/models"
	"encoding/json"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) repositories.Payment {
	return PaymentRepositoryImpl{
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

func (t PaymentRepositoryImpl) UpdatePaymentByNewEvent(orderId int64, status models.PaymentStatus, data paymentgateway.EventData) (*models.Payment, *models.PaymentEvent, error) {
	var exists bool

	results := t.db.Model(&models.Order{}).
		Select("count(*) > 0").
		Where("id = ?", orderId).
		Find(&exists)

	if results.Error != nil {
		slog.Info("order not exist", "err", results.Error)
		return nil, nil, results.Error
	}

	ctx := context.Background()

	payment, err := gorm.G[models.Payment](t.db).Where("order_id = ?", orderId).First(ctx)

	if err != nil {
		slog.Info("payment not exist", "err", err)
	}

	eventPayloadBytes, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal event data", "err", err)
		return nil, nil, err
	}
	rawPayload := json.RawMessage(eventPayloadBytes)

	paymentEvent := &models.PaymentEvent{
		PaymentID:         payment.ID,
		ServiceProviderID: payment.ServiceProviderID,
		Status:            status,
		EventPayload:      &rawPayload,
		CreatedAt:         time.Now(),
	}

	t.db.Model(&models.PaymentEvent{}).Create(paymentEvent)
	t.db.Model(&models.Payment{}).Where("id = ?", payment.ID).Update("status", status).Update("last_event_at", time.Now())

	return &payment, paymentEvent, nil
}
