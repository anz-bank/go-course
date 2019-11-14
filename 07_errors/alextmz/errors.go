package main

import "fmt"

type Error struct {
	Message string
	Code    int
}

const (
	ErrInvalidRequest = 400
	ErrNotFound       = 404
	ErrNegativeValue  = 499
)

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.errorCodeDescription(), e.Message)
}

func (e Error) errorCodeDescription() string {
	switch e.Code {
	case 0:
		panic("internal error: got no error code at ErrorCodeDescription")
	case 400:
		return "invalid request error"
	case 404:
		return "not found error"
	case 499:
		return "negative value error"
	default:
		return fmt.Sprintf("error %d", e.Code)
	}
}

func NewError(i int, m string) Error {
	var e Error
	e.Code = i
	e.Message = m
	return e
}

func NewErrorf(i int, f string, m ...interface{}) Error {
	var e Error
	e.Code = i
	e.Message = fmt.Sprintf(f, m...)
	return e
}
