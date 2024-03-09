package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nollidnosnhoj/vgpx/internal/config"
	"github.com/nollidnosnhoj/vgpx/internal/controllers"
	"github.com/nollidnosnhoj/vgpx/internal/router"
	"github.com/nollidnosnhoj/vgpx/internal/server"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cctx, cancel := context.WithCancel(context.Background())

	cfg := config.NewConfig()

	uploadStorage, err := storage.NewS3Storage(cctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter()

	homeController := controllers.NewHomeController()
	homeController.RegisterRoutes(router)
	uploadController := controllers.NewUploadController(uploadStorage)
	uploadController.RegisterRoutes(router)

	s := server.NewServer(router)

	// start server gorountine
	go s.Start(cctx)

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
