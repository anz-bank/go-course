package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestStringer(t *testing.T) {

	type test struct {
		ipInput  IPAddr
		expected string
	}

	tests := []test{
		{ipInput: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
		{ipInput: IPAddr{8, 8, 8, 8}, expected: "8.8.8.8"},
		{ipInput: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		{ipInput: IPAddr{192, 140, 123, 122}, expected: "192.140.123.122"},
	}

	for _, test := range tests {
		actual := test.ipInput.String()
		if actual != test.expected {
			t.Errorf("Unexpected output in \nexpected: %q\nactual: %q", test.expected, actual)
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
