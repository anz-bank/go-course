package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrCodeString(t *testing.T) {
	assert.Equal(t, "invalid input: %v", ErrInvalid.String())
	assert.Equal(t, "not found: %v", ErrNotFound.String())

	var ErrFoo ErrCode = 900
	assert.Equal(t, "error occurred", ErrFoo.String())
}

func TestError(t *testing.T) {
	err := NewError(ErrNotFound, "error message")
	assert.Equal(t, "not found: error message", err.Error())
}
