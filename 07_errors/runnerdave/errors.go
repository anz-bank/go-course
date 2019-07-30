package main

import "fmt"

// Error codes
const (
	// ErrUnknown is used when an unknown error occurred
	ErrUnknown uint16 = iota
	// ErrInvalidValue is used when the value for the puppy is negative
	ErrInvalidValue
	// ErrIDNotFound is used when attempting to read a non-existing entry
	ErrIDNotFound
)

// Error struct to identify errors in Puppy store
type Error struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Errorf creates a new Error with formatting
func Errorf(code uint16, format string, args ...interface{}) *Error {
	return &Error{code, fmt.Sprintf(format, args...)}
}

func validateValue(value float32) error {
	if value < 0 {
		return Errorf(ErrInvalidValue, "puppy has invalid value (%f)", value)
	}
	return nil
}
