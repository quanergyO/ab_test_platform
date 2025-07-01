package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/quanergyO/ab_test_platform/internal/handler"
	"github.com/quanergyO/ab_test_platform/internal/repository"
	"github.com/quanergyO/ab_test_platform/internal/repository/postgres"
	"github.com/quanergyO/ab_test_platform/internal/service"
	"github.com/quanergyO/ab_test_platform/server"
	"github.com/spf13/viper"
)

func main() {

	db, err := postgres.NewDB(postgres.Config{
		Username: "postgres",
		Host:     "db",
		Port:     "5432",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection", slog.Any("error", err))
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	serv := new(server.Server)
	go func() {
		if err := serv.Run("8080", handlers.InitRoutes()); err != nil {
			slog.Error("Error: failed to start server on port:", viper.GetString("port"), err.Error())
			os.Exit(1)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = serv.ShutDown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down:", err)
		os.Exit(1)
	}

	if err = db.Close(); err != nil {
		slog.Error("error occured on close db connection:", err)
		os.Exit(1)
	}
	slog.Info("Server shutting down")
}
