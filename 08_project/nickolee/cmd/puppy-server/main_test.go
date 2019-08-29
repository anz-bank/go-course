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

	expected := `Puppy with ID 1 has been created
Retrieved puppy: &{1 The Hound Of Baskerville 12300.9}
Update puppy operation result: <nil>
Delete puppy operation result: <nil>
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}
