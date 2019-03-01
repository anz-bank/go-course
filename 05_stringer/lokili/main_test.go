package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testSet = []struct {
	input  IPAddr
	output string
}{
	{IPAddr{192, 168, 31, 1}, "192.168.31.1"},
	{IPAddr{255, 255, 255, 0}, "255.255.255.0"},
}

func TestMainOutput(t *testing.T) {
	r := require.New(t)
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `127.0.0.1`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestStringer(t *testing.T) {
	r := require.New(t)
	for _, data := range testSet {
		r.Equalf(data.output, data.input.String(), "Unexpected output in main()")
	}
}
