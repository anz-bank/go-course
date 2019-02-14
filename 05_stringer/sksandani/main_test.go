package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	inputString  IPAddr
	outputString string
}{
	{IPAddr{127, 1, 2, 7}, "127.1.2.7"},
	{IPAddr{}, "0.0.0.0"},
}

func TestIpAddressMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := string("127.0.0.1\n")
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestIpAddress(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	for _, tt := range tests {
		buf.Truncate(0)
		fmt.Fprint(out, tt.inputString)
		actual := buf.String()
		r.Equal(tt.outputString, actual)
	}
}
