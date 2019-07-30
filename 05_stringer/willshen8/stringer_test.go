package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "127.0.0.1\n"
	actual := buf.String()

	assert.Equal(t, expected, actual)
}

var tests = []struct {
	name     string
	input    IPAddr
	expected string
}{
	{name: "Unspecified IP Address", input: IPAddr{}, expected: "0.0.0.0"},
	{name: "Empty IP Address", input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
	{name: "One Octet IP Address", input: IPAddr{127}, expected: "127.0.0.0"},
	{name: "Two Octets IP Address", input: IPAddr{127, 0}, expected: "127.0.0.0"},
	{name: "Three Octets IP Address", input: IPAddr{127, 0, 1}, expected: "127.0.1.0"},
	{name: "Broadcast IP Address", input: IPAddr{255, 255, 255, 255}, expected: "255.255.255.255"},
}

func TestStringer(t *testing.T) {

	for i := range tests {
		i := i
		t.Run(tests[i].name, func(t *testing.T) {
			assert.Equal(t, tests[i].expected, tests[i].input.String())
		})
	}
}
