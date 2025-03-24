package repository

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/eubide/get-quote/internal/core/domain"
	"github.com/eubide/get-quote/internal/core/ports"
)

// FileRepository implements the QuoteRepository interface using the file system
type FileRepository struct {
	config ports.ConfigProvider
	rng    *rand.Rand
}

// NewFileRepository creates a new FileRepository instance
func NewFileRepository(config ports.ConfigProvider) *FileRepository {
	// Create a new random number generator with a source seeded from the current time
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &FileRepository{
		config: config,
		rng:    rng,
	}
}

// GetRandomQuote retrieves a random quote from a file
func (r *FileRepository) GetRandomQuote(filename string) (*domain.Quote, error) {
	// Validate that the parameter has the correct extension
	if !strings.HasSuffix(filename, r.config.GetDefaultExtension()) {
		filename += r.config.GetDefaultExtension() // Add extension if not provided
	}

	// Construct the full path to the file in the files directory
	quotesFile := filepath.Join(r.config.GetFilesBaseDir(), filename)

	// Check if the file exists
	if _, err := os.Stat(quotesFile); os.IsNotExist(err) {
		return nil, domain.NewFileNotFoundError(quotesFile)
	}

	// Open the file
	file, err := os.Open(quotesFile)
	if err != nil {
		return nil, domain.NewFileOpenError(err)
	}
	defer file.Close()

	// Read all lines from the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, domain.NewFileOpenError(err)
	}

	// Check if there are any lines
	if len(lines) == 0 {
		return domain.NewQuote(""), nil
	}

	// Select a random line using the local random number generator
	randomIndex := r.rng.Intn(len(lines))
	randomLine := lines[randomIndex]

	return domain.NewQuote(randomLine), nil
}
