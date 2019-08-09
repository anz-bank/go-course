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
Update puppy result: true
Result of deleting puppy: true
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}
