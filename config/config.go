package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Http     *HttpConfig
	Database *DatabaseConfig
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Http:     LoadHttpConfig(),
		Database: LoadConfig(),
	}
}
