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
	expected := "Puppy ID 0 is 1280.5Puppy ID 0 is 1340.5"
	actual := buf.String()
	assert.Equal(expected, actual)
}
