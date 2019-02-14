package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected := fmt.Sprint("127.0.0.1")
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var cases = []struct {
	in  IPAddr
	out string
}{
	{IPAddr{10, 0, 0, 0}, "10.0.0.0"},
	{IPAddr{100, 100, 100, 0}, "100.100.100.0"},
	{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
	{IPAddr{110, 52, 32, 31}, "110.52.32.31"},
	{IPAddr{1, 2}, "1.2.0.0"},
	{IPAddr{}, "0.0.0.0"},
}

func TestInputs(t *testing.T) {
	//Given
	assert := assert.New(t)
	for _, e := range cases {
		res := IPAddr.String(e.in)
		assert.Equal(e.out, res)
	}
}
