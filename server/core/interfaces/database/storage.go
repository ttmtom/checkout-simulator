package database

import "crypto-checkout-simulator/server/core/interfaces/database/repositories"

/**
* Storage Interface defines the methods for accessing different repositories.
* This allows for a separation of concerns and makes it easier to test and maintain the code.
* Also it provide a way to switch between different storage implementations without changing the rest of the application.
 */

type Storage interface {
	GetPaymentRepository() repositories.Payment
}
