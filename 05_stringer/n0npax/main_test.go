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
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}

func TestStringify(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"localhost":            {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		"google primary dns":   {input: IPAddr{8, 8, 8, 8}, expected: "8.8.8.8"},
		"google secondary dns": {input: IPAddr{8, 8, 4, 4}, expected: "8.8.4.4"},
		"CloudFare":            {input: IPAddr{1, 1, 1, 1}, expected: "1.1.1.1"},
		"mask 24":              {input: IPAddr{255, 255, 255, 0}, expected: "255.255.255.0"},
		"mask 32":              {input: IPAddr{255, 255, 255, 255}, expected: "255.255.255.255"},
		"mask 0":               {input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
		"anz":                  {input: IPAddr{202, 2, 59, 40}, expected: "202.2.59.40"},
	}
	for name, test := range testCases {
		t.Run(name, helperTestStringify(test.input, test.expected))
	}
}

func helperTestStringify(input IPAddr, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := input.String()
		if actual != expected {
			t.Errorf("Expected: %q - Actual: %q", expected, actual)
		}
	}
}
