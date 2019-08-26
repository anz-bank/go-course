package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "Puppy(1) added to store\n"
	actual := buf.String()

	assert.Equal(t, expected, actual)
}
