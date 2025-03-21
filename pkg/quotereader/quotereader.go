package quotereader

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// Interface for random number generation to make testing easier
type RandGenerator interface {
	Intn(n int) int
}

// QuoteReader manages reading random quotes from a file
type QuoteReader struct {
	quotes []string
	rng    RandGenerator // Changed to interface type
}

// NewQuoteReader creates a new QuoteReader from the given file path
func NewQuoteReader(filePath string) (*QuoteReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			quotes = append(quotes, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Initialize random number generator with current time
	source := rand.NewSource(time.Now().UnixNano())

	return &QuoteReader{
		quotes: quotes,
		rng:    rand.New(source),
	}, nil
}

// GetRandomQuote returns a random quote from the loaded quotes
func (qr *QuoteReader) GetRandomQuote() string {
	if len(qr.quotes) == 0 {
		return ""
	}

	randomIndex := qr.rng.Intn(len(qr.quotes))
	return qr.quotes[randomIndex]
}
