package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	input        IPAddr
	outputString string
}{
	{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	{IPAddr{1, 1}, "1.1.0.0"},
	{IPAddr{}, "0.0.0.0"}}

func TestString(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	for _, tt := range tests {
		buf.Reset()
		fmt.Fprint(out, tt.input)
		actual := buf.String()
		r.EqualValues(tt.outputString, actual)
	}
}

func TestStringerMainOutput(t *testing.T) {
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
