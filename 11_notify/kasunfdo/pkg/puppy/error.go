package puppy

import "fmt"

type ErrCode uint32

type Error struct {
	Message string
	Code    ErrCode
	Cause   error
}

const (
	ErrInvalid   ErrCode = 400
	ErrNotFound  ErrCode = 404
	ErrBadFormat ErrCode = 406
	ErrInternal  ErrCode = 500
)

func (e ErrCode) String() string {
	switch e {
	case ErrInvalid:
		return "invalid input"
	case ErrNotFound:
		return "not found"
	case ErrBadFormat:
		return "bad data format"
	case ErrInternal:
		return "internal error"
	default:
		return "unknown error"
	}
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("%s: %s", e.Code, e.Message)
	if e.Cause == nil {
		return msg
	}
	return fmt.Sprintf("%s\n\t%s", msg, e.Cause)
}

func ErrorEf(code ErrCode, cause error, format string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
		Code:    code,
		Cause:   cause,
	}
}
