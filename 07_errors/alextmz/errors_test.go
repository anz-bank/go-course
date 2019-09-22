package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownError(t *testing.T) {
	e := NewError(ErrInvalidRequest, "test invalid request")
	assert.Equal(t, "invalid request error: test invalid request", e.Error())

	e = NewError(ErrNotFound, "test not found")
	assert.Equal(t, "not found error: test not found", e.Error())

	e = NewError(ErrNegativeValue, "test negative value")
	assert.Equal(t, "negative value error: test negative value", e.Error())

	e = Error{}
	assert.PanicsWithValue(t,
		"internal error: got no error code at ErrorCodeDescription",
		func() { e.errorCodeDescription() })
	e = NewError(999, "test error 999")
	assert.Equal(t, "error 999: test error 999", e.Error())
}

func TestNewError(t *testing.T) {
	e1 := NewError(999, "error 999")
	var e2 Error
	e2.Code = 999
	e2.Message = "error 999"
	assert.Equal(t, e2, e1)
}

func TestNewErrorf(t *testing.T) {
	e1 := NewErrorf(999, "error %d", 999)
	var e2 Error
	e2.Code = 999
	e2.Message = "error 999"
	assert.Equal(t, e2, e1)
}
