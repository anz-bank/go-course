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

	expected := `Puppy with ID %d has been created

	
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}
