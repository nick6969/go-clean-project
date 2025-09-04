package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nick6969/go-clean-project/internal/application"
	"github.com/nick6969/go-clean-project/internal/config"
	"github.com/nick6969/go-clean-project/internal/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := application.New(config)
	if err != nil {
		log.Fatalf("failed to create application: %v", err)
	}

	server, err := http.NewServer(app)
	if err != nil {
		app.Logger.Warn(context.Background(), "Failed to create server", "error", err)
		log.Fatalf("Failed to create server: %v", err)
	}

	server.Start()

	systemShutdownHandle()

	server.Shutdown()
}

func systemShutdownHandle() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
