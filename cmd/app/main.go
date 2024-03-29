package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/container"
	"github.com/nollidnosnhoj/kopalol/internal/db"
	"github.com/nollidnosnhoj/kopalol/internal/router"
	"github.com/nollidnosnhoj/kopalol/internal/server"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cctx, cancel := context.WithCancel(context.Background())
	config, err := config.NewConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(&config)

	if err != nil {
		log.Fatal(err)
	}

	storage, err := storage.NewStorage(cctx, storage.S3StorageType, &config)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.Default()

	container := &container.Container{
		Storage: storage,
		Db:      db,
		Logger:  logger,
		Config:  &config,
	}

	router := router.NewRouter(container)

	appServer := server.NewServer(router, container.Logger)
	go appServer.Start(cctx)

	// wait for signal to initiate shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	sig := <-quit

	log.Println("Shutting down...")
	cancel()

	if sig == syscall.SIGUSR1 {
		os.Exit(1)
	}
}
