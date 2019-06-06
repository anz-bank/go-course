package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestStringify(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"localhost": {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		"google":    {input: IPAddr{8, 8, 8, 8}, expected: "8.8.8.8"},
		"blank":     {input: IPAddr{}, expected: "0.0.0.0"},
		"quater":    {input: IPAddr{127}, expected: "127.0.0.0"},
		"half":      {input: IPAddr{127, 1}, expected: "127.1.0.0"},
		"3quaters":  {input: IPAddr{127, 1, 1}, expected: "127.1.1.0"},
	}
	for name, test := range testCases {
		input, expected := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			actual := input.String()
			if expected != actual {
				t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
			}
		})
	}
}

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
