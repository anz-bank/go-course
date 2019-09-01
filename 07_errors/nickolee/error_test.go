package puppystorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := Error{
		Message: "Dalmations",
		Code:    101,
	}

	errMsg := err.Error()
	assert.Equal(t, "PuppyStoreError 101: Dalmations", errMsg)
}
