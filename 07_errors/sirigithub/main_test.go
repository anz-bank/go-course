package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "Read puppy from Mapstore with ID: 0\nRead puppy from SyncStore with ID: 0\n"
	actual := buf.String()
	assert.Equal(t, expected, actual, "Expected and actual values should be the same.")
}
