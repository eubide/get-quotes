package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eubide/get-quote/internal/adapters/primary/cli"
	"github.com/eubide/get-quote/internal/adapters/secondary/config"
	"github.com/eubide/get-quote/internal/adapters/secondary/repository"
	"github.com/eubide/get-quote/internal/app/services"
)

func main() {
	// Parse command-line flags
	configFile := flag.String("c", "", "Path to configuration file")
	flag.Parse()

	// Update os.Args for the CLI handler (removing the -c flag and its value if present)
	if *configFile != "" {
		// Create a new slice with the program name
		newArgs := []string{os.Args[0]}
		
		// Add all arguments except -c and its value
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-c" {
				// Skip the -c flag and its value
				i++
				continue
			} else if len(os.Args[i]) > 1 && os.Args[i][0] == '-' && os.Args[i][1] == 'c' {
				// Skip -c flag combined with value (e.g. -cfile.yaml)
				continue
			}
			newArgs = append(newArgs, os.Args[i])
		}
		
		// Replace os.Args with the filtered version
		os.Args = newArgs
	}

	// Initialize configuration based on the defined order of preference
	cfg, err := config.NewYAMLConfigWithOrder(*configFile)
	if err != nil {
		fmt.Printf("Warning: Using default configuration: %v\n", err)
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
