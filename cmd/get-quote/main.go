package main

import (
	"fmt"
	"os"

	"github.com/eubide/get-quote/internal/adapters/primary/cli"
	"github.com/eubide/get-quote/internal/adapters/secondary/config"
	"github.com/eubide/get-quote/internal/adapters/secondary/repository"
	"github.com/eubide/get-quote/internal/app/services"
)

func main() {
	// Initialize configuration
	configPath := ".get-quote.yaml"
	cfg, err := config.NewYAMLConfig(configPath)
	if err != nil {
		fmt.Printf("Warning: Could not load config file: %v\n", err)
	}

	// Initialize repository
	repo := repository.NewFileRepository(cfg)

	// Initialize service
	service := services.NewQuoteService(repo, cfg)

	// Initialize CLI handler
	handler := cli.NewCLIHandler(service)

	// Execute the CLI handler
	if err := handler.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
