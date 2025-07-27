package main

import (
	"crypto-checkout-simulator/config"
	"crypto-checkout-simulator/pkg/logger"
	"crypto-checkout-simulator/pkg/utils"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	logger.Init()
	godotenv.Load()

	actionList := []string{"up", "down"}
	argType := "up"
	if len(os.Args) > 1 {
		if !utils.Contains(actionList, os.Args[1]) {
			slog.Error("Invalid action type", "arg", os.Args)
			os.Exit(1)
		}

		argType = os.Args[1]
	}
	slog.Info(argType)

	pgConfig := config.LoadConfig()

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.DBName,
		"sslmode=disable",
	)

	sourceURL := "file://server/adapter/storage/pg/migrations"

	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)
	if err != nil {
		slog.Error("Error on DB connection", "err", err)
		os.Exit(1)
	}

	if argType == "up" {
		err = m.Up()
	} else if argType == "down" {
		err = m.Steps(-1)
	}

	if err != nil {
		if err.Error() == "no change" {
			slog.Info(fmt.Sprintf("Migration %s completed with no changes", argType))
		} else {
			slog.Error(fmt.Sprintf("Error on migration %s", argType), "err", err)
			os.Exit(1)
		}
	}
	slog.Info(fmt.Sprintf("Migration %s completed successfully", argType))
}
