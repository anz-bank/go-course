package puppy

// Error codes
const (
	// ErrInvalidInput is used when the incoming request is invalid
	ErrInvalidInput int = iota
	// ErrDuplicate is used when attempting to erroneously overwrite an existing entry
	ErrDuplicate
	// ErrNotFound is used when attempting to read a non-existing entry
	ErrNotFound
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
