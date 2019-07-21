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
		t.Errorf("actual: %s does not match expected: %s", actual, expected)
	}
}
func TestStringMethod(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"happy path":  {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		"one value":   {input: IPAddr{127}, expected: "127.0.0.0"},
		"two value":   {input: IPAddr{127, 0}, expected: "127.0.0.0"},
		"three value": {input: IPAddr{127, 0, 0}, expected: "127.0.0.0"},
		"max value":   {input: IPAddr{255, 255, 255, 255}, expected: "255.255.255.255"},
		"empty":       {input: IPAddr{}, expected: "0.0.0.0"},
	}
	for name, tc := range testCases {
		actual := tc.input.String()
		expected := tc.expected
		t.Run(name, func(t *testing.T) {
			if expected != actual {
				t.Errorf("got %q wanted %q", actual, expected)
			}
		})
	}
}
