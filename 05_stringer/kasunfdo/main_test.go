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

	expected := "127.0.0.1"
	actual := buf.String()

	assert.Equalf(t, expected, actual, "Unexpected output")
}

func TestIPAddrStringer(t *testing.T) {
	tests := map[string]struct {
		input    IPAddr
		expected string
	}{
		"Uninitialized IP": {
			input:    IPAddr{},
			expected: "0.0.0.0",
		},
		"Semi-initialized": {
			input:    IPAddr{1, 2},
			expected: "1.2.0.0",
		},
		"Non-routable IP": {
			input:    IPAddr{0, 0, 0, 0},
			expected: "0.0.0.0",
		},
		"Broadcast IP": {
			input:    IPAddr{255, 255, 255, 255},
			expected: "255.255.255.255",
		},
		"Localhost": {
			input:    IPAddr{127, 0, 0, 1},
			expected: "127.0.0.1",
		},
	}

	for name, test := range tests {
		input, expected := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			actual := input.String()
			assert.Equalf(t, expected, actual, "Unexpected output")
		})
	}
}
