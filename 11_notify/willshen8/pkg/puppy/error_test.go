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

func TestInternalError(t *testing.T) {
	internalErr := ErrCode(2)
	internalErrString := internalErr.String()
	assert.Equal(t, "Internal Error", internalErrString)
}
