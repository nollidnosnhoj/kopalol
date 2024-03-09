package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nollidnosnhoj/vgpx/internal/cache"
	"github.com/nollidnosnhoj/vgpx/internal/server"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
)

func main() {
	cctx, cancel := context.WithCancel(context.Background())

	// config := config.NewConfig()

	// // _, err := database.NewDatabase(config)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	localStorage, err := storage.NewLocalStorage("uploads")
	if err != nil {
		log.Fatal(err)
	}
	cache := cache.NewCache()
	s := server.NewServer(cache, localStorage)

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
