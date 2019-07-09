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
	expected := `First Created Puppy ID: 1
ReadPuppy: &{1 Jack Russell White and Brown 500}
Updated Puppy: &{1 Fox Terrier Black 1300}
Result of deleting puppy: true
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}
