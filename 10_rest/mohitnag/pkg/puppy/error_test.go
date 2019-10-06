package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert := assert.New(t)
	err := ErrorF(NotFound, "test error message")
	assert.Equal(NotFound, err.Code)
	assert.Equal("test error message", err.Message)
	assert.Equal("test error message :(code: puppy not found)", err.Error())
}
