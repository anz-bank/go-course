package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "127.0.0.1"
	got := buf.String()

	if expected != got {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
	}
}

func TestIPAddr(t *testing.T) {
	var cases = map[string]struct {
		input    IPAddr
		expected string
	}{
		"Empty": {
			input:    IPAddr{},
			expected: "0.0.0.0",
		},
		"Partial - one value": {
			input:    IPAddr{1},
			expected: "1.0.0.0",
		},
		"Partial - two values": {
			input:    IPAddr{1, 2},
			expected: "1.2.0.0",
		},
		"Zeros": {
			input:    IPAddr{0, 0, 0, 0},
			expected: "0.0.0.0",
		},
		"255s": {
			input:    IPAddr{255, 255, 255, 255},
			expected: "255.255.255.255",
		},
	}

	for name, c := range cases {
		got, expected := c.input.String(), c.expected
		t.Run(name, func(t *testing.T) {
			if got != expected {
				t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
			}
		})
	}
}
