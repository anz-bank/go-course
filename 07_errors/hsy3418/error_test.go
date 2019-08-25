package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidInputError(t *testing.T) {
	err := Errorf(ErrInvalidInput, "invalid input")
	assert.Equal(t, err.Code, ErrInvalidInput)
	errMessage := err.Error()
	assert.Equal(t, "Error code:0,Error:invalid input", errMessage)
}

func TestNotFoundError(t *testing.T) {
	err := Errorf(ErrNotFound, "error not found")
	assert.Equal(t, err.Code, ErrNotFound)
	errMessage := err.Error()
	assert.Equal(t, "Error code:2,Error:error not found", errMessage)
}

func TestDuplicateError(t *testing.T) {
	err := Errorf(ErrDuplicate, "duplicate value")
	assert.Equal(t, err.Code, ErrDuplicate)
	errMessage := err.Error()
	assert.Equal(t, "Error code:1,Error:duplicate value", errMessage)
}
