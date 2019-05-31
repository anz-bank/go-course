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
		t.Errorf("Unexpected output printed by main.\n(Expected) : %v\n(Actual)   : %v", expected, actual)
	}
}

func TestFib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	type test struct {
		input    int
		expected string
	}

	tests := []test{
		{-7, `"1\n-1\n2\n-3\n5\n-8\n13\n"`},
		{7, `"1\n1\n2\n3\n5\n8\n13\n"`},
		{-3, `"1\n-1\n2\n"`},
		{4, `"1\n1\n2\n3\n"`},
	}

	for _, test := range tests {
		fib(test.input)
		actual := strconv.Quote(buf.String())

		if test.expected != actual {
			t.Errorf("Unexpected output printed by fib.\n(Expected) : %v\n(Actual)   : %v", test.expected, actual)
		}

		buf.Reset()
	}
}
