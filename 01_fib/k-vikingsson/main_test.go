package main

import (
	"bytes"
	"testing"
)

var buf bytes.Buffer

var testCases = map[string]struct {
	input    int
	expected string
}{
	"Basic":    {input: 7, expected: "1\n1\n2\n3\n5\n8\n13\n"},
	"Nagetive": {input: -7, expected: "0\n1\n-1\n2\n-3\n5\n-8\n13\n"},
	"Zero":     {input: 0, expected: "0\n"},
}

func TestFib(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		expected := test.expected
		t.Run(name, func(t *testing.T) {
			out = &buf
			fib(input)
			actual := buf.String()
			if test.expected != actual {
				t.Errorf("Unexpected output in fib():\nexpected: %v\nactual: %v\n", expected, actual)
			}
			buf.Reset()
		})
	}
}

func TestMain(t *testing.T) {
	out = &buf
	main()
	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in main():\nexpected: %v\nactual: %v\n", expected, actual)
	}
}