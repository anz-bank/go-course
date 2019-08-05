package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Expected %v, got %v ", actual, expected)
	}
}

func TestFib(t *testing.T) {

	testCases := map[string]struct {
		input    int
		expected string
	}{
		"positive": {input: 7, expected: strconv.Quote(`1
1
2
3
5
8
13
`)},
		"negative": {input: -7, expected: strconv.Quote(`1
-1
2
-3
5
-8
13
`)},
		"0": {input: 0, expected: strconv.Quote(`0
`)},
	}

	for name, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		var buf bytes.Buffer
		out = &buf
		t.Run(name, func(t *testing.T) {
			fib(test.input)
			result := strconv.Quote(buf.String())
			if result != test.expected {
				t.Errorf("Expected %v, got %v ", test.expected, result)
			}
		})
	}
}
