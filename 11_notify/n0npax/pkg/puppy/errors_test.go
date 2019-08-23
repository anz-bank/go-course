package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	err := Errorf(ErrCodeInternal, "internal error")
	assert.Equal(t, ErrCodeInternal, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "internal error (500)", errMessage)
}
