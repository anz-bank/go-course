package puppy

import "fmt"

const (
	// Invalid is used when the puppy value is below 0
	Invalid = "invalid input"
	// NotFound is used when the puppy with given id is not found
	NotFound = "puppy not found"
	//Duplicate is used when the puppy id already exists
	Duplicate = "puppy already exists"
)

// Error is the authorisationservice package error with a code for comparison
type Error struct {
	Message string
	Code    string
}

// Error method implments the error interface for message retrieval
func (e *Error) Error() string {
	return fmt.Sprintf("%s :(code: %s)", e.Message, e.Code)
}

// ErrorF is a utility function creating an error with given code and message
func ErrorF(code, format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return &Error{Message: message, Code: code}
}
