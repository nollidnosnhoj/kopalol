package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nollidnosnhoj/vgpx/internal/config"
	"github.com/nollidnosnhoj/vgpx/internal/controllers"
	"github.com/nollidnosnhoj/vgpx/internal/database"
	"github.com/nollidnosnhoj/vgpx/internal/queries"
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbConn, err := database.Open(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	queries := queries.New(dbConn)
	uploadStorage, err := storage.NewS3Storage(cctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter(logger)

	homeController := controllers.NewHomeController()
	homeController.RegisterRoutes(router)
	uploadController := controllers.NewUploadController(queries, uploadStorage, logger)
	uploadController.RegisterRoutes(router)
	filesController := controllers.NewFilesController(queries, uploadStorage, logger)
	filesController.RegisterRoutes(router)

	appServer := server.NewServer(router, logger)
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
