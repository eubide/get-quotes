package ports

import "github.com/eubide/get-quote/internal/core/domain"

// QuoteService defines the interface for the application service
type QuoteService interface {
	// GetRandomQuote retrieves a random quote from a specified file
	GetRandomQuote(filename string) (*domain.Quote, error)
}
