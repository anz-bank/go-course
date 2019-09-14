package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	StrErrGeneric = "error, generic"
)

func TestError_Error(t *testing.T) {
	assert.Condition(t, func() bool {
		return NewError(0).Message == StrErrGeneric
	})
	assert.Condition(t, func() bool {
		return NewError(1).Message == StrErrGeneric
	})
	assert.Condition(t, func() bool {
		return NewError(ErrGeneric).Message == StrErrGeneric
	})
	assert.Condition(t, func() bool {
		return NewError(8).Message == "error, ID being deleted does not exist"
	})
	assert.Condition(t, func() bool {
		return NewError(ErrIDBeingDeletedDoesNotExist).Message == "error, ID being deleted does not exist"
	})
}
