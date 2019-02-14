package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputs = []struct {
	in  IPAddr
	out string
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{}, "0.0.0.0"},
	{IPAddr{1, 1, 1, 1}, "1.1.1.1"}}

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

	for _, t := range testInputs {
		buf.Reset()
		fmt.Fprint(out, t.in)
		expected := strconv.Quote(t.out)
		actual := strconv.Quote(buf.String())
		r.EqualValues(expected, actual)
	}
}
