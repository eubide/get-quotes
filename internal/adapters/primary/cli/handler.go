package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eubide/get-quote/internal/core/domain"
	"github.com/eubide/get-quote/internal/core/ports"
)

// CLIHandler handles command-line interface interactions
type CLIHandler struct {
	quoteService ports.QuoteService
	config       ports.ConfigProvider
}

// NewCLIHandler creates a new CLIHandler instance
func NewCLIHandler(quoteService ports.QuoteService) *CLIHandler {
	// Extract the config from the quote service
	// This is a safe type assertion since we know QuoteServiceImpl has a config field
	config, ok := quoteService.(interface{ GetConfig() ports.ConfigProvider })
	var configProvider ports.ConfigProvider
	if ok {
		configProvider = config.GetConfig()
	}

	return &CLIHandler{
		quoteService: quoteService,
		config:       configProvider,
	}
}

// Execute processes the command-line arguments and executes the appropriate action
func (h *CLIHandler) Execute() error {
	// Check if a file parameter was provided
	if len(os.Args) < 2 {
		return h.handleMissingParameter()
	}

	// Get the file parameter
	fileName := os.Args[1]

	// Handle special flags
	if fileName == "-h" || fileName == "--help" {
		return h.printHelp()
	}

	if fileName == "-l" || fileName == "--list" {
		return h.listAvailableFiles()
	}

	// Get a random quote
	quote, err := h.quoteService.GetRandomQuote(fileName)
	if err != nil {
		// Handle specific error types
		switch err.(type) {
		case *domain.DomainError:
			// For file not found errors, show available files
			return fmt.Errorf("error: %v\n\nAvailable files:\n%s", err, h.getAvailableFilesString())
		default:
			return err
		}
	}

	// Print the random quote without a newline
	fmt.Print(quote.Text)

	return nil
}

// handleMissingParameter provides a helpful error message when the filename is missing
func (h *CLIHandler) handleMissingParameter() error {
	programName := filepath.Base(os.Args[0])
	message := fmt.Sprintf("Usage: %s [options] <filename>\n\n", programName)
	message += "Options:\n"
	message += "  -h, --help    Show this help message\n"
	message += "  -l, --list    List available quote files\n\n"
	message += "Available files:\n"
	message += h.getAvailableFilesString()
	
	return fmt.Errorf(message)
}

// printHelp prints the usage information
func (h *CLIHandler) printHelp() error {
	programName := filepath.Base(os.Args[0])
	
	fmt.Printf("Get Quote - A simple tool for extracting random quotes from text files\n\n")
	fmt.Printf("Usage: %s [options] <filename>\n\n", programName)
	fmt.Printf("Options:\n")
	fmt.Printf("  -c <file>     Specify a configuration file path\n")
	fmt.Printf("  -h, --help    Show this help message\n")
	fmt.Printf("  -l, --list    List available quote files\n\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("  %s quotes     Extract a random quote from quotes.lst\n", programName)
	fmt.Printf("  %s citas      Extract a random quote from citas.lst\n\n", programName)
	fmt.Printf("Available files:\n")
	fmt.Print(h.getAvailableFilesString())
	
	return nil
}

// listAvailableFiles lists all available quote files
func (h *CLIHandler) listAvailableFiles() error {
	fmt.Println("Available quote files:")
	fmt.Print(h.getAvailableFilesString())
	return nil
}

// getAvailableFilesString returns a string listing all available quote files
func (h *CLIHandler) getAvailableFilesString() string {
	if h.config == nil {
		return "  Unable to determine available files: configuration not available\n"
	}

	baseDir := h.config.GetFilesBaseDir()
	ext := h.config.GetDefaultExtension()
	
	// Check if directory exists
	_, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		return fmt.Sprintf("  Directory '%s' not found\n", baseDir)
	}

	// Read directory contents
	files, err := os.ReadDir(baseDir)
	if err != nil {
		return fmt.Sprintf("  Error reading directory '%s': %v\n", baseDir, err)
	}

	result := ""
	found := false

	// List all files with the appropriate extension
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		
		name := file.Name()
		if filepath.Ext(name) == ext {
			// Show filename without extension
			baseName := name[:len(name)-len(ext)]
			result += fmt.Sprintf("  %s\n", baseName)
			found = true
		}
	}

	if !found {
		result = fmt.Sprintf("  No files with '%s' extension found in '%s'\n", ext, baseDir)
	}

	return result
}