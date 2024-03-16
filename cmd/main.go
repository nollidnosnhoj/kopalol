package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	router := router.NewRouter()
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	homeController := controllers.NewHomeController()
	homeController.RegisterRoutes(router)
	uploadController := controllers.NewUploadController(queries, uploadStorage, logger)
	uploadController.RegisterRoutes(router)

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
