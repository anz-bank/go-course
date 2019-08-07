package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "127.0.0.1"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestStringMethod(t *testing.T) {
	type test struct {
		name     string
		input    IPAddr
		expected string
	}

	tests := []test{
		{name: "all local IPv4 addresses", input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
		{name: "google dns/a lucky number :)", input: IPAddr{8, 8, 8, 8}, expected: "8.8.8.8"},
		{name: "localhost/loopback", input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		{name: "wikipedia", input: IPAddr{208, 80, 154, 224}, expected: "208.80.154.224"},
		{name: "facebook", input: IPAddr{31, 13, 71, 36}, expected: "31.13.71.36"},
		{name: "nil value", input: IPAddr{}, expected: "0.0.0.0"}, // defaults to the nil value of a byte
		{name: "missing", input: IPAddr{192, 150}, expected: "192.150.0.0"},
	}

	for _, testCase := range tests {
		actual := testCase.input.String()
		if actual != testCase.expected {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
