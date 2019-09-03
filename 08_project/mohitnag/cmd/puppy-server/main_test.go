package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "Mapstore: white\nSyncstore: white\n"
	actual := buf.String()
	assert.Equal(expected, actual)
}
