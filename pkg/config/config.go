package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds application configuration
type Config struct {
	// Base directory for files
	FilesBaseDir string `yaml:"filesBaseDir"`
	// Default file extension
	DefaultExtension string `yaml:"defaultExtension"`
	// Error messages
	ErrorMessages struct {
		FileNotFound     string `yaml:"fileNotFound"`
		FileOpenError    string `yaml:"fileOpenError"`
		MissingParameter string `yaml:"missingParameter"`
	} `yaml:"errorMessages"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		FilesBaseDir:     filepath.Join("src", "files"),
		DefaultExtension: ".lst",
		ErrorMessages: struct {
			FileNotFound     string `yaml:"fileNotFound"`
			FileOpenError    string `yaml:"fileOpenError"`
			MissingParameter string `yaml:"missingParameter"`
		}{
			FileNotFound:     "El archivo %s no existe",
			FileOpenError:    "Error al abrir el archivo: %v",
			MissingParameter: "Uso: %s <nombre_fichero>\nDebe proporcionar un nombre de fichero %s",
		},
	}
}

// FindConfigFile searches for a config file in multiple paths
func FindConfigFile(filename string) (string, bool) {
	// Search paths in order with specific names:
	// 1. ./get-quote.yaml (current directory, visible)
	// 2. $HOME/.config/get-quote/get-quote.yaml
	// 3. $HOME/.get-quote.yaml (invisible)

	// Check current directory with visible name
	visibleConfig := "get-quote.yaml"
	if _, err := os.Stat(visibleConfig); err == nil {
		return visibleConfig, true
	}

	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err == nil {
		// Check $HOME/.config/get-quote/get-quote.yaml
		configDir := filepath.Join(homeDir, ".config", "get-quote")
		configPath := filepath.Join(configDir, "get-quote.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, true
		}

		// Check $HOME/.get-quote.yaml (invisible file)
		homePath := filepath.Join(homeDir, ".get-quote.yaml")
		if _, err := os.Stat(homePath); err == nil {
			return homePath, true
		}
	}

	return "", false
}

// LoadConfig loads configuration from a file
func LoadConfig(configName string) (*Config, error) {
	// Find config file in search paths
	configPath, found := FindConfigFile(configName)

	// If no config file is found, create one in $HOME/.config/get-quote/
	if !found {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return DefaultConfig(), err
		}

		// Create config directory if it doesn't exist
		configDir := filepath.Join(homeDir, ".config", "get-quote")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return DefaultConfig(), err
		}

		// Create new config file
		configPath = filepath.Join(configDir, "get-quote.yaml")
		config := DefaultConfig()

		// Marshal config to YAML
		data, err := yaml.Marshal(config)
		if err != nil {
			return DefaultConfig(), err
		}

		// Write config file
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return DefaultConfig(), err
		}

		return config, nil
	}

	// Read and parse config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return DefaultConfig(), err
	}

	config := DefaultConfig()
	if err := yaml.Unmarshal(data, config); err != nil {
		return DefaultConfig(), err
	}

	return config, nil
}
