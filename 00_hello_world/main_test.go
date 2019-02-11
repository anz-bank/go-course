package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "0\n1\n1\n2\n3\n5\n8\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
