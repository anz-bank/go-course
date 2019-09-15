package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownError(t *testing.T) {
	randomErrorCode := ErrCode(100000)
	randomErrorCodeString := randomErrorCode.String()
	assert.Equal(t, "Unknown error", randomErrorCodeString)
}
