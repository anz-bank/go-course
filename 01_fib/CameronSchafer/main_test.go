package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

//main() test
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

//fib(int) test
func TestFibOutput(t *testing.T) {
	//tempBuf is used to reset the output buffer for each test
	var buf, tempBuf bytes.Buffer
	out = &buf

	//test cases with descriptions.
	testCases := []struct {
		description string
		input       int
		expected    string
	}{
		{description: "fib 7", input: 7,
			expected: strconv.Quote("1\n1\n2\n3\n5\n8\n13\n"),
		},
		{description: "fib 1", input: 1,
			expected: strconv.Quote("1\n"),
		},
		{description: "fib -7", input: -7,
			expected: strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n"),
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			fib(test.input)
			result := strconv.Quote(buf.String())
			if result != test.expected {
				t.Errorf("%v\nexpected %v, got %v", test.description, test.expected, result)
			}
			buf = tempBuf //reset the buffer for the next test.
		})
	}
}

//calculateNormalFib(int) []int test
func TestNormalFibOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description string
		input       int
		expected    []int
	}{
		{description: "calculateNormalFib 7", input: 7,
			expected: []int{1, 1, 2, 3, 5, 8, 13},
		},
		{description: "calculateNormalFib 1", input: 1,
			expected: []int{1},
		},
		{description: "calculateNormalFib 1", input: 0,
			expected: nil,
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			result := calculateNormalFib(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Unexpected output in %q\nexpected: %q\nactual: %q", test.description, test.expected, result)
			}
		})
	}
}

//NegaFibOutput test
func TestNegaFibOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description string
		input       int
		expected    []int
	}{
		{description: "calculateNegaFib 7", input: 7,
			expected: []int{1, -1, 2, -3, 5, -8, 13},
		},
		{description: "calculateNegaFib ", input: 2,
			expected: []int{1, -1},
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			result := calculateNegaFib(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Unexpected output in %q\nexpected: %q\nactual: %q", test.description, test.expected, result)
			}
		})
	}
}

//CalcNextInSequence test
func TestCalcNextInSequence(t *testing.T) {
	expected := 5
	actual := calcNextInSequence(1, 4)

	if actual != expected {
		t.Errorf("Unexpected output in calcNextInSequence(int,int)\nexpected: %q\nactual: %q", expected, actual)
	}
}

//PrintSequence test
func TestPrintSequence(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	printFibSequence([]int{1, 2, 3, 6, 23})

	expected := strconv.Quote("1\n2\n3\n6\n23\n")
	actual := strconv.Quote(buf.String())

	if actual != expected {
		t.Errorf("Unexpected output in printFibSequence(int)\nexpected: %q\nactual: %q", expected, actual)
	}
}
