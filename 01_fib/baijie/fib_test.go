package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected %v, actual %v", expected, actual)
	}
}

func TestFib1(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 0, expected: ""},
		{input: 1, expected: "1\n"},
		{input: 7, expected: "1\n1\n2\n3\n5\n8\n13\n"},
		{input: -7, expected: "1\n-1\n2\n-3\n5\n-8\n13\n"},
	}

	var buf bytes.Buffer
	out = &buf

	for _, test := range tests {
		fib(test.input)
		actual := buf.String()

		if test.expected != actual {
			t.Fatalf("expected: %v, got: %v", test.expected, actual)
		}
		buf.Reset()

	}
}
