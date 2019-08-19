package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownError(t *testing.T) {
	a := assert.New(t)

	error := NewError(123)
	a.Equal(error, NewError(Unknown))
}

func TestError(t *testing.T) {
	a := assert.New(t)

	error := Error{1, "message"}
	formattedError := error.Error()
	a.Equal(formattedError, "[1] message")
}
