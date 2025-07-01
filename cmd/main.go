package main

import (
	"log/slog"
	"os"

	"github.com/quanergyO/ab_test_platform/internal/handler"
	"github.com/quanergyO/ab_test_platform/internal/repository"
	"github.com/quanergyO/ab_test_platform/internal/repository/postgres"
	"github.com/quanergyO/ab_test_platform/internal/service"
	"github.com/spf13/viper"
)

func main() {
	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		slog.Error("Error: failed to init db connection", slog.Any("error", err))
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	_ = handler.NewHandler(services)

}
