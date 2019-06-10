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
	testCases := []struct {
		description string
		ip          IPAddr
		expected    string
	}{
		{"happy path", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"one value", IPAddr{127}, "127.0.0.0"},
		{"two value", IPAddr{127, 0}, "127.0.0.0"},
		{"three value", IPAddr{127, 0, 0}, "127.0.0.0"},
		{"empty", IPAddr{}, "0.0.0.0"},
	}
	for _, tc := range testCases {
		actual := tc.ip
		expected := tc.expected
		t.Run(tc.description, func(t *testing.T) {
			actual := actual.String()
			if expected != actual {
				t.Errorf("got %q wanted %q", actual, expected)
			}
		})
	}
}
