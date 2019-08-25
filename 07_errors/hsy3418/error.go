package main

import "fmt"

type Error struct {
	Message string
	Code    ErrorCode
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code:%d,Error:%s", e.Code, e.Message)

}

// Errorf creates a new Error with formatting
func Errorf(code ErrorCode, format string, args ...interface{}) *Error {
	return ErrorEf(code, format, args...)
}

// ErrorEf creates a new Error with causing error and formatting
func ErrorEf(code ErrorCode, format string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
		Code:    code,
	}
}

//ErrCode defines the kind of error
type ErrorCode uint8

//Error codes
const (
	// ErrInvalidInput for invalid user input
	ErrInvalidInput ErrorCode = iota
	// ErrDuplicate is used when attempting to create an already existing entry
	ErrDuplicate
	// ErrNotFound is used when trying to access a non-existing entry
	ErrNotFound
	// ErrNotFound is used when initialisation or setup fails
)
