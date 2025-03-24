package ports

// ConfigProvider defines the interface for accessing configuration
type ConfigProvider interface {
	// GetFilesBaseDir returns the base directory for quote files
	GetFilesBaseDir() string

	// GetDefaultExtension returns the default file extension for quote files
	GetDefaultExtension() string

	// GetErrorMessages returns the error messages configuration
	GetErrorMessages() ErrorMessages
}

// ErrorMessages defines the structure for error message templates
type ErrorMessages struct {
	FileNotFound     string
	FileOpenError    string
	MissingParameter string
}
