package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of fib(7) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote(`1
1
2
3
5
8
13`)
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestFibFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    []int
	}{
		{"Test zero fib to return zero", 0, []int{0}},
		{"Test positive fib one to return one", 1, []int{1}},
		{"Test positive fib two to return the correct values", 2, []int{1, 1}},
		{"Test positive fib to return the correct values", 7, []int{1, 1, 2, 3, 5, 8, 13}},
		{"Test negative fib one to return one", -1, []int{1}},
		{"Test negative fib two to return the correct values", -2, []int{1, -1}},
		{"Test negative fib to return the correct values", -7, []int{1, -1, 2, -3, 5, -8, 13}},
	}

	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected

		t.Run(testCase.description, func(t *testing.T) {
			actual := fib(input)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Unexpected output, expected: %d, actual: %d", expected, actual)
			}
		})
	}
}
