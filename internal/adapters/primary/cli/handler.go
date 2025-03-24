package cli

import (
	"fmt"
	"os"

	"github.com/eubide/get-quote/internal/core/ports"
)

// CLIHandler handles command-line interface interactions
type CLIHandler struct {
	quoteService ports.QuoteService
}

// NewCLIHandler creates a new CLIHandler instance
func NewCLIHandler(quoteService ports.QuoteService) *CLIHandler {
	return &CLIHandler{
		quoteService: quoteService,
	}
}

// Execute processes the command-line arguments and executes the appropriate action
func (h *CLIHandler) Execute() error {
	// Check if a file parameter was provided
	if len(os.Args) < 2 {
		return fmt.Errorf("missing file parameter")
	}

	// Get the file parameter
	fileName := os.Args[1]

	// Get a random quote
	quote, err := h.quoteService.GetRandomQuote(fileName)
	if err != nil {
		return err
	}

	// Print the random quote without a newline
	fmt.Print(quote.Text)

	return nil
}
