package ports

import "github.com/eubide/get-quote/internal/core/domain"

// QuoteRepository defines the interface for accessing quotes
type QuoteRepository interface {
	// GetRandomQuote retrieves a random quote from the repository
	GetRandomQuote(filename string) (*domain.Quote, error)
}
