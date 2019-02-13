package main

import (
	"bytes"
	"fmt"
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
	expected := `127.0.0.1`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var s = []struct {
	input  IPAddr
	output string
}{

	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{10, 10, 40, 1}, "10.10.40.1"},
	{IPAddr{7, 20, 210}, "7.20.210.0"},
	{IPAddr{7, 20}, "7.20.0.0"},
	{IPAddr{0}, "0.0.0.0"},
	{IPAddr{}, "0.0.0.0"},
}

func TestIPAddrInterface(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	for _, v := range s {
		expected := fmt.Sprint(v.input)
		actual := v.output
		r.Equalf(expected, actual, "Unexpected output in IPAddr()")
	}
}
