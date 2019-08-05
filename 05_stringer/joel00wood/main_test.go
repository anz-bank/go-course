package main

import (
	"testing"

	output "github.com/joel00wood/test-helpers/capture"
)

func TestMain(t *testing.T) {
	expected := "127.0.0.1\n"
	actual := output.CaptureOutput(func() { main() })
	if expected != actual {
		t.Errorf("Unexpected result in main(), expected=%q, got=%q",
			expected, actual)
	}
}

func TestIPAddr(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"Standard input": {
			IPAddr{127, 0, 0, 1},
			"127.0.0.1",
		},
		"googleDNS": {
			IPAddr{8, 8, 8, 8},
			"8.8.8.8",
		},
		"Empty input": {
			IPAddr{},
			"0.0.0.0",
		},
		"Partially initialised": {
			IPAddr{4, 20},
			"4.20.0.0",
		},
	}

	for name, test := range testCases {
		actual := test.input.String()
		if test.expected != actual {
			t.Errorf("Unexpected response for %v, expected=%q, got=%q",
				name, test.expected, actual)
		}
	}
}
