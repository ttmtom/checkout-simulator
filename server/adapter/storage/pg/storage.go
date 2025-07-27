package pg

import (
	"crypto-checkout-simulator/config"
	"crypto-checkout-simulator/server/adapter/storage/pg/impl_repositories"
	"crypto-checkout-simulator/server/core/interfaces/database"
	"crypto-checkout-simulator/server/core/interfaces/database/repositories"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresql struct {
	payment *impl_repositories.PaymentRepositoryImpl
	order   *impl_repositories.OrderRepositoryImpl
}

func New(config *config.DatabaseConfig) (database.Storage, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	PostgresDb, dbOpenErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbOpenErr != nil {
		return nil, dbOpenErr
	}

	PaymentRepo := impl_repositories.NewPaymentRepoImpl(PostgresDb)
	OrderRepo := impl_repositories.NewOrderRepoImpl(PostgresDb)

	return &Postgresql{
		payment: PaymentRepo,
		order:   OrderRepo,
	}, nil
}

func (p *Postgresql) GetPaymentRepository() repositories.Payment {
	return p.payment
}

func (p *Postgresql) GetOrderRepository() repositories.Order {
	return p.order
}
