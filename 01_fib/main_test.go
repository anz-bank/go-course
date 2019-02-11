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
	expected := "-8 -5 -3 -2 -1 -1 0 1 1 2 3 5 8"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
