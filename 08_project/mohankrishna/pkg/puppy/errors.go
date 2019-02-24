package puppy

// Error codes
const (
	ErrInvalidInput int = iota
	ErrDuplicate
	ErrNotFound
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
