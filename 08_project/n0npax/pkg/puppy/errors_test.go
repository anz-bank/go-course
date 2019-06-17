package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotFound(t *testing.T) {
	err := ErrNotFound("test not found")
	assert.Equal(t, ErrNotFoundCode, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code: 404, message : test not found", errMessage)
}

func TestErrorInvalidInput(t *testing.T) {
	err := ErrInvalidInput("test invalid input")
	assert.Equal(t, ErrInvalidInputCode, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code: 422, message : test invalid input", errMessage)
}

func TestErrInternalError(t *testing.T) {
	err := ErrInternalError("internal error")
	assert.Equal(t, ErrInternalErrorCode, err.Code)
	errMessage := err.Error()
	assert.Equal(t, "Error code: 500, message : internal error", errMessage)
}
