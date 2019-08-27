package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotFound(t *testing.T) {
	err := Errorf(ErrCodeNotFound, "test not found")
	assert.Equal(t, ErrCodeNotFound, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "test not found (404)", errMessage)
}

func TestErrorInvalidInput(t *testing.T) {
	err := Errorf(ErrCodeInvalidInput, "test invalid input")
	assert.Equal(t, ErrCodeInvalidInput, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "test invalid input (400)", errMessage)
}

func TestErrorf(t *testing.T) {
	err := Errorf(ErrCodeInternal, "internal error")
	assert.Equal(t, ErrCodeInternal, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "internal error (500)", errMessage)
}
