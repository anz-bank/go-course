package main

import "fmt"

type Error struct {
	Message string
	Code    int
}

const (
	ErrInvalidValue = 400
	ErrIDNotFound   = 404
)

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}
func NewError(errCode int, msg string) *Error {
	return &Error{msg, errCode}
}
