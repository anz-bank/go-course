package puppy

import "fmt"

type Error struct {
	Message string  `json:"message"`
	Code    ErrCode `json:"code"`
}

type ErrCode uint32

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}

// Errorf creates a new Error with formatting
func Errorf(code ErrCode, message string, args ...interface{}) *Error {
	return &Error{
		fmt.Sprintf(message, args...),
		code,
	}
}

const (
	// ErrInvalidInput is used for any user input error
	ErrInvalidInput ErrCode = iota
	// ErrNotFound is used for puppy ID does not exist
	ErrNotFound
	// ErrInternalError is used for any internal error
	ErrInternalError
)

func (e ErrCode) String() string {
	switch e {
	case ErrInvalidInput:
		return "Invalid input, ensure ID and JSON are valid"
	case ErrNotFound:
		return "The puppy ID does not exist"
	case ErrInternalError:
		return "Internal Error"
	}
	return "Unknown error"
}
