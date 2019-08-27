package puppystorer

import "fmt"

// Error (our custom error type) wraps errors with code, message and error itself - what does this even mean??
type Error struct {
	Message string
	Code    int
}

// Error() method lets us satisfy error interface for our custom error type
func (e *Error) Error() string {
	return fmt.Sprintf("PuppyStoreError %d: %s", e.Code, e.Message)
}

// Error codes
const (
	ErrNegativePuppyID = 400
	ErrPuppyNotFound   = 404
)
