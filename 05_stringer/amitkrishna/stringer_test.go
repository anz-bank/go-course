package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData = []struct {
	i        IPAddr // input
	expected string // expected result
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{0, 0, 0, 1}, "0.0.0.1"},
	{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
	{IPAddr{0, 0, 0, 1}, "0.0.0.1"},
}

func TestNumeronymOutput(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	for _, tt := range testData {
		buf.Reset()
		fmt.Fprint(out, tt.i)
		actual := buf.String()
		r.Equalf(tt.expected, actual, "Unexpected output in main()")
	}
}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "127.0.0.1"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
