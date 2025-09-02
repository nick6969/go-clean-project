package main

import (
	"log"

	"github.com/nick6969/go-clean-project/internal/application"
	"github.com/nick6969/go-clean-project/internal/config"
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

	// 這裡先只把 application 印出，後續會使用到
	log.Println(app)
}
