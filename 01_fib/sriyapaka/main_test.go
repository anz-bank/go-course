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
	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestAbsOutput(t *testing.T) {

	tests := map[string]struct {
		input    int
		expected int
	}{
		"abs(7)":  {input: 7, expected: 7},
		"abs(-7)": {input: -7, expected: 7},
		"abs(1)":  {input: 1, expected: 1},
		"abs(-1)": {input: -1, expected: 1},
		"abs(0)":  {input: 0, expected: 0},
	}

	for name, test := range tests {

		var buf bytes.Buffer
		out = &buf
		test := test

		t.Run(name, func(t *testing.T) {

			actual := abs(test.input)
			if test.expected != actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, actual)
			}
		})
	}
}

func TestFibSeries(t *testing.T) {

	tests := map[string]struct {
		input    int
		expected string
	}{
		"fib(7)":  {input: 7, expected: strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")},
		"fib(-7)": {input: -7, expected: strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n")},
		"fib(1)":  {input: 1, expected: strconv.Quote("1\n")},
		"fib(-1)": {input: -1, expected: strconv.Quote("1\n")},
	}

	for name, test := range tests {

		var buf bytes.Buffer
		out = &buf
		test := test

		t.Run(name, func(t *testing.T) {

			fib(test.input)
			actual := strconv.Quote(buf.String())
			if test.expected != actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, actual)
			}
		})
	}
}
