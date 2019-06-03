package main

import (
	"bytes"
	"strconv"
	"testing"

	assertion "github.com/stretchr/testify/assert"
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
		"empty":                {input: IPAddr{}, expected: "0.0.0.0"},
		"half-empty":           {input: IPAddr{10, 10}, expected: "10.10.0.0"},
	}
	for name, test := range testCases {
		input, expected := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			assert := assertion.New(t)
			actual := input.String()
			assert.Equalf(expected, actual, "Testcase %s failed")
		})
	}
}
