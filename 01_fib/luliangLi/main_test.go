package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFibOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	testCases := map[string]struct {
		input int
		want  string
	}{
		"7":  {input: 7, want: "1\n1\n2\n3\n5\n8\n13\n"},
		"-7": {input: -7, want: "1\n-1\n2\n-3\n5\n-8\n13\n"},
		"0":  {input: 0, want: "0\n"},
		"1":  {input: 1, want: "1\n"},
		"-1": {input: -1, want: "1\n"},
		"2":  {input: 2, want: "1\n1\n"},
		"-2": {input: -2, want: "1\n-1\n"},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		test := test

		t.Run(name, func(t *testing.T) {
			fib(test.input)

			expected := strconv.Quote(test.want)
			actual := strconv.Quote(buf.String())

			if expected != actual {
				t.Errorf("runing n: %v, expected %v, got %v", test.input, expected, actual)
			}
			buf.Reset()
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("runing n: %v, expected %v, got %v", "7", expected, actual)
	}
}
