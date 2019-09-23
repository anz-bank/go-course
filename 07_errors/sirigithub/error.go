package main

import "fmt"

type Error struct {
	Message string
	Code    int
}

const (
	ErrInvalidValue   = 400
	ErrIDNotFound     = 404
	ErrMarshallData   = 400
	ErrUrmarshallData = 400
	ErrDatabaseConn   = 500
	ErrDatabaseWrite  = 500
	ErrDatabaseRead   = 500
	ErrDatabseDelete  = 500
)

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}
func NewError(errCode int, msg string) *Error {
	return &Error{msg, errCode}
}
