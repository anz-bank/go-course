package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in  IPAddr
	out string
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{}, "0.0.0.0"},
	{IPAddr{1, 1, 1, 1}, "1.1.1.1"},
}

func TestMainOutput(t *testing.T) {
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

func TestStringOutput(t *testing.T) {
	r := require.New(t)

	for _, tt := range tests {
		r.Equal(tt.out, tt.in.String())
	}
}
