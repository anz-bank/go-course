package puppy

import (
	"fmt"
)

// Error codes
const (
	ErrInvalidInputCode  = 400
	ErrNotFoundCode      = 404
	ErrInternalErrorCode = 500
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
func Errorf(code int, format string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
		Code:    code,
	}
}
