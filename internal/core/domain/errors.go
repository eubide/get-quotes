package domain

import "fmt"

// DomainError represents a domain-specific error
type DomainError struct {
	Message string
}

// Error implements the error interface
func (e *DomainError) Error() string {
	return e.Message
}

// NewFileNotFoundError creates a new error for when a file is not found
func NewFileNotFoundError(filename string) *DomainError {
	return &DomainError{
		Message: fmt.Sprintf("The file %s does not exist", filename),
	}
}

// NewFileOpenError creates a new error for when a file cannot be opened
func NewFileOpenError(err error) *DomainError {
	return &DomainError{
		Message: fmt.Sprintf("Error opening the file: %v", err),
	}
}

// NewMissingParameterError creates a new error for when a required parameter is missing
func NewMissingParameterError(programName, extension string) *DomainError {
	return &DomainError{
		Message: fmt.Sprintf("Usage: %s <file_name>\nYou must provide a file name %s", programName, extension),
	}
}
