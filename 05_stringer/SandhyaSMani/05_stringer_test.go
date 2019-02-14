package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in  IPAddress
	out string
}{
	{IPAddress{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddress{}, "0.0.0.0"},
	{IPAddress{192, 0, 2, 1}, "192.0.2.1"},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestStringOutput(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	for _, testVal := range tests {
		buf.Reset()
		fmt.Fprint(out, testVal.in)
		expected := strconv.Quote(testVal.out)
		actual := strconv.Quote(buf.String())
		r.EqualValues(expected, actual)
	}
}
