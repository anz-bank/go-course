package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testMatrix = []struct {
	input          IPAddr
	expectedOutput string
}{
	{IPAddr{255, 123, 39, 40}, "255.123.39.40"},
	{IPAddr{255, 0, 0, 0}, "255.0.0.0"},
	{IPAddr{0, 123, 39, 40}, "0.123.39.40"},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `12.23.23.24`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestStringer(t *testing.T) {
	// Given
	r := require.New(t)

	for _, testData := range testMatrix {
		r.EqualValues(testData.expectedOutput, testData.input.String())
	}
}
