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
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := testCase.input.String()
			if testCase.expected != actual {
				t.Errorf("Unexpected output, expected: %s, actual: %s", testCase.expected, actual)
			}
		})
	}
}
