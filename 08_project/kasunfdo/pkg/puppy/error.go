package puppy

import "fmt"

type ErrCode uint32

type Error struct {
	Message string
	Code    ErrCode
}

const (
	ErrInvalid  ErrCode = 400
	ErrNotFound ErrCode = 404
)

func (e ErrCode) String() string {
	switch e {
	case ErrInvalid:
		return "invalid input: %v"
	case ErrNotFound:
		return "not found: %v"
	default:
		return "error occurred"
	}
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code ErrCode, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(code.String(), args...),
		Code:    code,
	}
}
