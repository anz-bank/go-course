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

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	main()
	// Then
	expected := `192.0.0.1
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestStringer(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		output := fmt.Sprintf("%v", tt.input)
		r.Equal(tt.outputString, output)
	}
}
