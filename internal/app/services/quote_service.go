package services

import (
	"os"

	"github.com/eubide/get-quote/internal/core/domain"
	"github.com/eubide/get-quote/internal/core/ports"
)

// QuoteServiceImpl implements the QuoteService interface
type QuoteServiceImpl struct {
	repository ports.QuoteRepository
	config     ports.ConfigProvider
}

// NewQuoteService creates a new QuoteServiceImpl instance
func NewQuoteService(repository ports.QuoteRepository, config ports.ConfigProvider) *QuoteServiceImpl {
	return &QuoteServiceImpl{
		repository: repository,
		config:     config,
	}
}

// GetRandomQuote retrieves a random quote from a specified file
func (s *QuoteServiceImpl) GetRandomQuote(filename string) (*domain.Quote, error) {
	// Check if a file parameter was provided
	if filename == "" {
		return nil, domain.NewMissingParameterError(os.Args[0], s.config.GetDefaultExtension())
	}

	// Get a random quote from the repository
	return s.repository.GetRandomQuote(filename)
}
