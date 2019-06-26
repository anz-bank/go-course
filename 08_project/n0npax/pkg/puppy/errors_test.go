package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	err := Errorf(ErrInternalErrorCode, "internal error")
	assert.Equal(t, ErrInternalErrorCode, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code: 500, message : internal error", errMessage)
}
