package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/eubide/get-quote/internal/core/ports"
	"gopkg.in/yaml.v3"
)

// YAMLConfig implements the ConfigProvider interface using YAML files
type YAMLConfig struct {
	FilesBaseDir   string `yaml:"filesBaseDir"`
	DefaultExtension string `yaml:"defaultExtension"`
	ErrorMessages  struct {
		FileNotFound     string `yaml:"fileNotFound"`
		FileOpenError    string `yaml:"fileOpenError"`
		MissingParameter string `yaml:"missingParameter"`
	} `yaml:"errorMessages"`
}

// NewYAMLConfig creates a new YAMLConfig instance
func NewYAMLConfig(configPath string) (*YAMLConfig, error) {
	// Default configuration
	config := &YAMLConfig{
		FilesBaseDir:   "src/files",
		DefaultExtension: ".lst",
		ErrorMessages: struct {
			FileNotFound     string `yaml:"fileNotFound"`
			FileOpenError    string `yaml:"fileOpenError"`
			MissingParameter string `yaml:"missingParameter"`
		}{
			FileNotFound:     "The file %s does not exist",
			FileOpenError:    "Error opening the file: %v",
			MissingParameter: "Usage: %s <file_name>\nYou must provide a file name %s",
		},
	}

	// Try to load configuration from file
	if err := config.loadFromFile(configPath); err != nil {
		// Try to load from home directory
		homeDir, err := os.UserHomeDir()
		if err == nil {
			homePath := filepath.Join(homeDir, ".get-quote.yaml")
			if err := config.loadFromFile(homePath); err != nil {
				// Try to load from config directory
				configDir := filepath.Join(homeDir, ".config", "get-quote", "get-quote.yaml")
				_ = config.loadFromFile(configDir) // Ignore error, use defaults
			}
		}
	}

	return config, nil
}

// loadFromFile loads configuration from a YAML file
func (c *YAMLConfig) loadFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("could not parse config file: %w", err)
	}

	return nil
}

// GetFilesBaseDir returns the base directory for quote files
func (c *YAMLConfig) GetFilesBaseDir() string {
	return c.FilesBaseDir
}

// GetDefaultExtension returns the default file extension for quote files
func (c *YAMLConfig) GetDefaultExtension() string {
	return c.DefaultExtension
}

// GetErrorMessages returns the error messages configuration
func (c *YAMLConfig) GetErrorMessages() ports.ErrorMessages {
	return ports.ErrorMessages{
		FileNotFound:     c.ErrorMessages.FileNotFound,
		FileOpenError:    c.ErrorMessages.FileOpenError,
		MissingParameter: c.ErrorMessages.MissingParameter,
	}
}
