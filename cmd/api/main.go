package main

import (
	"log"

	"github.com/nick6969/go-clean-project/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println(config)
}
