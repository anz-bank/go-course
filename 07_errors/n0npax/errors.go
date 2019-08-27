package main

import (
	"fmt"
)

// Error codes
const (
	ErrCodeInvalidInput = 400
	ErrCodeNotFound     = 404
	ErrCodeInternal     = 500
)

// Error wrapps errors with code, message and error itself
type Error struct {
	Message string
	Code    int
}

// Error returns error as a string
func (e *Error) Error() string {
	return fmt.Sprintf("%v (%v)", e.Message, e.Code)
}

// Errorf creates a new Error with formatting
func Errorf(code int, format string, args ...interface{}) *Error {
	return &Error{
		fmt.Sprintf(format, args...),
		code,
	}
}
