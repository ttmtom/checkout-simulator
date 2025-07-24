package config

import "os"

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig() *DatabaseConfig {
	dbName := os.Getenv("POSTGRES_DATABASE_NAME")
	if dbName == "" {
		dbName = "postgres"
	}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}

	return &DatabaseConfig{
		DBName:   dbName,
		Host:     host,
		Port:     port,
		Password: password,
		User:     user,
	}
}
