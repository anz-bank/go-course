package main

import (
	"testing"

	output "github.com/joel00wood/test-helpers/capture"
)

func TestFibCalc(t *testing.T) {
	testCases := map[string]struct {
		input, expected float64
	}{
		"input 20": {20, 6765},
		"input 3":  {3, 2},
		"input 2":  {2, 1},
		"input 1":  {1, 1},
		"input 0":  {0, 0},
	}

	for name, test := range testCases {
		actual := fibCalc(test.input)
		if test.expected != actual {
			t.Errorf("tests[%s] fibCalc(%f) wrong. expected=%f, got=%f",
				name, test.input, test.expected, actual)
		}
	}
}

func TestFibN(t *testing.T) {
	testCases := map[string]struct {
		input, expected float64
	}{
		"input 20":  {20, 6765},
		"input 3":   {3, 2},
		"input 2":   {2, 1},
		"input 1":   {1, 1},
		"input 0":   {0, 0},
		"input -1":  {-1, 1},
		"input -2":  {-2, -1},
		"input -3":  {-3, 2},
		"input -20": {-20, -6765},
	}

	for name, test := range testCases {
		actual := fibN(test.input)
		if test.expected != actual {
			t.Errorf("tests[%s] fibN(%f) wrong. expected=%f, got=%f",
				name, test.input, test.expected, actual)
		}
	}
}

func TestFib(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected string
	}{
		"input 7":  {7, "0\n1\n1\n2\n3\n5\n8\n13\n"},
		"input 3":  {3, "0\n1\n1\n2\n"},
		"input 2":  {2, "0\n1\n1\n"},
		"input 1":  {1, "0\n1\n"},
		"input 0":  {0, "0\n"},
		"input -1": {-1, "0\n1\n"},
		"input -2": {-2, "0\n1\n-1\n"},
		"input -3": {-3, "0\n1\n-1\n2\n"},
		"input -7": {-7, "0\n1\n-1\n2\n-3\n5\n-8\n13\n"},
	}
	for name, test := range testCases {
		input := test.input
		actual := output.CaptureOutput(func() { fib(input) })
		if test.expected != actual {
			t.Errorf("tests[%s] fib(%d) wrong. expected=%q, got=%q",
				name, test.input, test.expected, actual)
		}
	}
}

// test main() which is essentially fib(7)
func TestMain(t *testing.T) {
	expected := "0\n1\n1\n2\n3\n5\n8\n13\n"
	actual := output.CaptureOutput(func() { main() })
	if expected != actual {
		t.Errorf("Unexpected response for input main(){fib(7)}\nExpected: %q\nActual: %q",
			expected, actual)
	}
}
