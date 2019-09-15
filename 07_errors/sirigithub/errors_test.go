package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidValue(t *testing.T) {
	err := NewError(ErrInvalidValue, "test invalid value")
	assert.Equal(t, ErrInvalidValue, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code 400: test invalid value", errMessage)
}

func TestIDNotFound(t *testing.T) {
	err := NewError(ErrIDNotFound, "test id not found")
	assert.Equal(t, ErrIDNotFound, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code 404: test id not found", errMessage)
}
