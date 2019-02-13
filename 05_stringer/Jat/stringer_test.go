package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringerOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "127.0.0.1\n"

	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
