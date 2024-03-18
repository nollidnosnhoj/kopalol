package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nollidnosnhoj/kopalol/config"
	"github.com/nollidnosnhoj/kopalol/router"
	"github.com/nollidnosnhoj/kopalol/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cctx, cancel := context.WithCancel(context.Background())

	container, err := config.NewContainer(cctx)

	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter(container)

	appServer := server.NewServer(router, container.Logger())
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
