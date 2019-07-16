package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
		t.Errorf("\nActual: %s\nExpected: %s", actual, expected)
	}
}

func TestIPAddr(t *testing.T) {
	tests := map[string]struct {
		input    IPAddr
		expected string
	}{
		"localhost": {
			IPAddr{127, 0, 0, 1},
			"127.0.0.1",
		},
		"255.255.255.255": {
			IPAddr{255, 255, 255, 255},
			"255.255.255.255",
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := test.input.String()
			if actual != test.expected {
				t.Errorf("result: %v, expected %v", actual, test.expected)
			}
		})
	}
}
