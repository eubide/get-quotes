package randomline

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/eubide/get-quote/pkg/config"
)

// Define an interface for random number generation
type randInterface interface {
	Intn(n int) int
}

// Make the rand function mockable for testing
var randNewFunc = func(source rand.Source) randInterface {
	return rand.New(source)
}

// GetRandomLine returns a random line from a file specified in the configuration
func GetRandomLine(fileName string, configPath string) (string, error) {
	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return "", fmt.Errorf("error loading configuration: %w", err)
	}

	// Process file name
	if !strings.HasSuffix(fileName, cfg.DefaultExtension) {
		fileName += cfg.DefaultExtension // Add extension if not provided
	}

	// Construct the full path to the file
	filePath := filepath.Join(cfg.FilesBaseDir, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf(cfg.ErrorMessages.FileNotFound, filePath)
	}

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(cfg.ErrorMessages.FileOpenError, err)
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Check if we have any lines
	if len(lines) == 0 {
		return "", fmt.Errorf("file is empty")
	}

	// Select a random line
	source := rand.NewSource(time.Now().UnixNano())
	rng := randNewFunc(source)
	randomIndex := rng.Intn(len(lines))

	// Return the line without any potential newline characters
	return strings.TrimSpace(lines[randomIndex]), nil
}

// GetRandomLines returns a specified number of random lines from a text file
// This function is kept for backward compatibility
func GetRandomLines(filePath string, count int) ([]string, error) {
	// Validate count
	if count <= 0 {
		return nil, fmt.Errorf("count must be greater than 0")
	}

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Count total lines in file
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Check if we have enough lines
	if lineCount == 0 {
		return nil, fmt.Errorf("file is empty")
	}
	if count > lineCount {
		count = lineCount
	}

	// Reset file pointer to beginning
	if _, err := file.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("failed to reset file: %w", err)
	}

	// Select random lines
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	selectedIndices := make(map[int]bool)
	for len(selectedIndices) < count {
		selectedIndices[rng.Intn(lineCount)] = true
	}

	// Read the file again and collect selected lines
	scanner = bufio.NewScanner(file)
	result := make([]string, 0, count)
	currentIndex := 0
	for scanner.Scan() {
		if selectedIndices[currentIndex] {
			// Trim any potential newline characters from the line
			result = append(result, strings.TrimSpace(scanner.Text()))
		}
		currentIndex++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return result, nil
}
