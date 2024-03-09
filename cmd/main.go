package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/nollidnosnhoj/vgpx/internal/cache"
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

	cache := cache.NewCache(cache.CacheSettings{
		Expiration:      24 * time.Hour,     // expire after 24 hours
		CleanupInterval: 24 * 7 * time.Hour, // cleanup every 7 days
	})

	router := router.NewRouter()

	homeController := controllers.NewHomeController()
	homeController.RegisterRoutes(router)
	imageController := controllers.NewImageController(cache, uploadStorage)
	imageController.RegisterRoutes(router)
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
