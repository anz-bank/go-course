package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrCodeString(t *testing.T) {
	assert.Equal(t, "invalid input", ErrInvalid.String())
	assert.Equal(t, "not found", ErrNotFound.String())
	assert.Equal(t, "internal error", ErrInternal.String())
	assert.Equal(t, "bad data format", ErrBadFormat.String())

	var ErrFoo ErrCode = 900
	assert.Equal(t, "unknown error", ErrFoo.String())
}

func TestError(t *testing.T) {
	err := ErrorEf(ErrNotFound, nil, "error message")
	assert.Equal(t, "not found: error message", err.Error())
}
