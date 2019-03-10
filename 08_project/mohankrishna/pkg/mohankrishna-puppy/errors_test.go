package puppy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorsInvalid(t *testing.T) {
	// Given
	r := require.New(t)
	err := &Error{ErrInvalidInput, "Error invalid input"}
	errMsg := fmt.Sprint(err)

	r.Equalf("Error invalid input", errMsg, "Unexpected error message")
	r.Equalf(err.Code, ErrInvalidInput, "Unexpected error code")
}

func TestErrorsDuplicate(t *testing.T) {
	// Given
	r := require.New(t)
	err := &Error{ErrDuplicate, "Duplicate, already exists"}
	errMsg := fmt.Sprint(err)

	r.Equalf("Duplicate, already exists", errMsg, "Unexpected error message")
	r.Equalf(err.Code, ErrDuplicate, "Unexpected error code")
}

func TestErrorsNotFound(t *testing.T) {
	// Given
	r := require.New(t)
	err := &Error{ErrNotFound, "Not found"}
	errMsg := fmt.Sprint(err)

	r.Equalf("Not found", errMsg, "Unexpected error message")
	r.Equalf(err.Code, ErrNotFound, "Unexpected error code")
}
