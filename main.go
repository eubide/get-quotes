package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/eubide/get-quote/pkg/config"
	"github.com/eubide/get-quote/pkg/quotereader"
)

func main() {
	// Load configuration from YAML file
	cfg, err := config.LoadConfig(".get-quote.yaml")
	if err != nil {
		// Just log the error, we'll use default config
		log.Printf("Warning: Could not load config file: %v", err)
	}

	// Check if a file parameter was provided
	if len(os.Args) < 2 {
		log.Fatalf(cfg.ErrorMessages.MissingParameter, os.Args[0], cfg.DefaultExtension)
	}

	// Get the file parameter
	fileParam := os.Args[1]

	// Validate that the parameter has the correct extension
	if !strings.HasSuffix(fileParam, cfg.DefaultExtension) {
		fileParam += cfg.DefaultExtension // Add extension if not provided
	}

	// Construct the full path to the file in the files directory
	quotesFile := filepath.Join(cfg.FilesBaseDir, fileParam)

	// Check if the file exists
	if _, err := os.Stat(quotesFile); os.IsNotExist(err) {
		log.Fatalf(cfg.ErrorMessages.FileNotFound, quotesFile)
	}

	reader, err := quotereader.NewQuoteReader(quotesFile)
	if err != nil {
		log.Fatalf(cfg.ErrorMessages.FileOpenError, err)
	}

	quote := reader.GetRandomQuote()
	fmt.Print(quote)
}
