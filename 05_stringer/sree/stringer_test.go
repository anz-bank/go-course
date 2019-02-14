package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in       IPAddr
	Expected string
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{127, 0, 0}, "127.0.0.0"},
	{IPAddr{127, 0}, "127.0.0.0"},
	{IPAddr{0}, "0.0.0.0"},
	{IPAddr{}, "0.0.0.0"},
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

func TestIPAddr(t *testing.T) {
	// Given
	assert := assert.New(t)

	for _, v := range tests {
		assert.Equalf(v.Expected, fmt.Sprint(v.in), "Unexpected output in IPAddr()")
	}
}
