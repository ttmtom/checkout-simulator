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
	Transaction *impl_repositories.TransactionRepositoryImpl
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

	TransactionRepo := impl_repositories.NewTransactionRepoImpl(PostgresDb)

	return &Postgresql{
		Transaction: TransactionRepo,
	}, nil
}

func (s *Postgresql) GetTransactionRepository() repositories.TransactionRepository {
	return s.Transaction
}
