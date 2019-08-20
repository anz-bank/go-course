package puppy

import "fmt"

type ErrorCode int

type Error struct {
	Code    ErrorCode
	Message string
}

const (
	NegativeValue ErrorCode = iota + 1001
	IDNotFound
	Unknown = 1999
)

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError creates a new error with the given enum
func NewError(code ErrorCode) error {
	switch code {
	case NegativeValue:
		return &Error{code, "Puppy value must be greater than 0"}
	case IDNotFound:
		return &Error{code, "Nonexistent Puppy ID"}
	default:
		return &Error{Unknown, "Unknown error"}
	}
}
