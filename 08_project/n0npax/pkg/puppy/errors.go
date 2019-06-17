package puppy

import (
	"fmt"
)

// Error codes
const (
	ErrInvalidInputCode  = 422
	ErrNotFoundCode      = 404
	ErrInternalErrorCode = 500
)

// Error descriptions
const (
	PuppyNotFoundMsg   = "Puppy with ID %v not found"
	InvalidInputMsg    = "Puppy value have to be positive number"
	DataDecodeErrorMsg = "Internal error. Could not cast stored data to puppy object"
	CorruptedIDMsg     = "ID is corrupted. Please ensure object ID matched provided ID"
)

// Error wrapps errors with code, message and error itself
type Error struct {
	Message string
	Code    int
}

// Error returns error as a string
func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %v, message : %v", e.Code, e.Message)
}

// Errorf creates a new Error with formatting
func Errorf(code int, msg string) *Error {
	return &Error{
		Message: msg,
		Code:    code,
	}
}

// ErrInvalidInput generates custom error for invalid input
func ErrInvalidInput(msg string) *Error {
	return Errorf(ErrInvalidInputCode, msg)
}

// ErrNotFound generates custom error for not found resource
func ErrNotFound(msg string) *Error {
	return Errorf(ErrNotFoundCode, msg)
}

// ErrInternalError generates custom error for internal error
func ErrInternalError(msg string) *Error {
	return Errorf(ErrInternalErrorCode, msg)
}
