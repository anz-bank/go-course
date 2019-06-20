package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of IPAddr({127, 0, 0, 1}) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote("127.0.0.1")
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestIpAddrDefaultToString(t *testing.T) {
	testCases := []struct {
		description string
		input       IPAddr
		expected    string
	}{
		{"Test IPAddr with all zeros in addresses", IPAddr{0, 0, 0, 0}, "0.0.0.0"},
		{"Test IPAddr with some values assigned", IPAddr{add1: 5, add4: 4}, "5.0.0.4"},
		{"Test IPAddr with no value", IPAddr{}, "0.0.0.0"},
		{"Test IPAddr with max values allowed on uint8", IPAddr{255, 255, 255, 255}, "255.255.255.255"},
	}

	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected

		t.Run(testCase.description, func(t *testing.T) {
			actual := input.String()
			if expected != actual {
				t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
			}
		})
	}
}
