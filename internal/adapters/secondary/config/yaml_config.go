package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
// This function is kept for backward compatibility
func NewYAMLConfig(configPath string) (*YAMLConfig, error) {
	return NewYAMLConfigWithOrder("")
}

// NewYAMLConfigWithOrder creates a new YAMLConfig instance respecting the following order of preference:
// 1. If specifiedPath is provided, use it
// 2. Look for get-quote.yaml in ~/.config/get-quote/
// 3. Look for .get-quote.yaml in user's home directory
// 4. Look for get-quote.yaml in the current directory
func NewYAMLConfigWithOrder(specifiedPath string) (*YAMLConfig, error) {
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

	var err error
	configLoaded := false

	// 1. If specifiedPath is provided, use it (highest priority)
	if specifiedPath != "" {
		if err = config.loadFromFile(specifiedPath); err == nil {
			configLoaded = true
		}
	}

	// Only proceed with other paths if no config loaded yet
	if !configLoaded {
		// Get the user's home directory
		homeDir, homeDirErr := os.UserHomeDir()
		
		// 2. Look for get-quote.yaml in ~/.config/get-quote/
		if !configLoaded && homeDirErr == nil {
			configDir := filepath.Join(homeDir, ".config", "get-quote", "get-quote.yaml")
			if err = config.loadFromFile(configDir); err == nil {
				configLoaded = true
			}
		}

		// 3. Look for .get-quote.yaml in the user's home directory
		if !configLoaded && homeDirErr == nil {
			homePath := filepath.Join(homeDir, ".get-quote.yaml")
			if err = config.loadFromFile(homePath); err == nil {
				configLoaded = true
			}
		}

		// 4. Look for get-quote.yaml in the current directory
		if !configLoaded {
			if err = config.loadFromFile("get-quote.yaml"); err == nil {
				configLoaded = true
			}
		}
	}

	// Return the config even if we couldn't load a file - it will use defaults
	if !configLoaded {
		return config, fmt.Errorf("could not load any configuration file, using defaults")
	}

	return config, nil
}

// loadFromFile loads configuration from a YAML file
func (c *YAMLConfig) loadFromFile(path string) error {
	// Expand tilde if present
	expandedPath, err := expandTilde(path)
	if err != nil {
		return fmt.Errorf("could not expand path: %w", err)
	}

	data, err := os.ReadFile(expandedPath)
	if err != nil {
		return fmt.Errorf("could not read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("could not parse config file: %w", err)
	}

	return nil
}

// expandTilde replaces the tilde (~) character with the user's home directory
func expandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}

	// Handle the case of "~" alone by using homeDir directly
	if path == "~" {
		return homeDir, nil
	}

	// Handle paths like "~/something" by removing the tilde and slash, 
	// then joining with home directory
	if len(path) > 1 && path[1] == '/' {
		return filepath.Join(homeDir, path[2:]), nil
	}

	// Handle paths like "~something" which shouldn't be expanded
	return path, nil
}

// GetFilesBaseDir returns the base directory for quote files
func (c *YAMLConfig) GetFilesBaseDir() string {
	// Try to expand tilde if present
	expandedPath, err := expandTilde(c.FilesBaseDir)
	if err != nil {
		// If there's an error, return the original path
		return c.FilesBaseDir
	}
	return expandedPath
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
