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

//	@title					Go Clean Architecture API
//	@version				0.0.1
//	@description 		This is Go Clean Architecture API
//	@termsOfService
//	@host 					go-clean-architecture.xxxx.com

//	@securityDefinitions.apikey Bearer
//	@in header
//	@name Authorization
//	@description Type "Bearer" followed by a space and JWT token.

//	@BasePath	/
//	@schemes	https

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

	err = app.Database.MigrateUp()
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
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
