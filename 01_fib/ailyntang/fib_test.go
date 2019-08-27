package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestUnexpectedInputs(t *testing.T) {
	errorString := strconv.Quote("fib.go only handles integers between 1 to 92 inclusive\n")

	tests := []struct {
		input    int
		expected string
	}{
		{0, errorString},
		{93, errorString},
		{183, errorString},
		{-5, errorString},
		{-205, errorString},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		out = &buf

		fib(test.input)

		expected := test.expected
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected response from fib(%v)", test.input)
			t.Errorf("expected: %v\nactual: %v", expected, actual)
		}
	}
}

func TestGenerateFibSequence(t *testing.T) {
	tests := []struct {
		input    int
		expected []int64
	}{
		{1, []int64{1}},
		{7, []int64{1, 1, 2, 3, 5, 8, 13}},
		{10, []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}

	for _, test := range tests {
		actual := genFib(test.input)
		expected := test.expected
		errorMessage := "Incorrect Fibonnaci sequence\nexpected: %v\nactual: %v"

		if len(expected) != len(actual) {
			t.Errorf(errorMessage, expected, actual)
		}
		for i := range expected {
			if expected[i] != actual[i] {
				t.Errorf(errorMessage, expected, actual)
			}
		}
	}
}

func TestGenFibUpperLimitOf92(t *testing.T) {
	indices := []int{40, 60, 91}
	expected := make([]int64, 92)
	expected[40] = 165580141
	expected[60] = 2504730781961
	expected[91] = 7540113804746346429

	actual := genFib(92)

	if len(expected) != len(actual) {
		t.Errorf("Lengths of slices are not equal")
		t.Errorf("expected: %v\nactual: %v", len(expected), len(actual))
	}

	for _, indexValue := range indices {
		if expected[indexValue] != actual[indexValue] {
			t.Errorf("Incorrect Fibonnaci number at index %v", indexValue)
			t.Errorf("expected: %v\nactual: %v", expected[indexValue], actual[indexValue])
		}
	}
}

func TestFibOutput(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, strconv.Quote("1\n")},
		{6, strconv.Quote("1\n1\n2\n3\n5\n8\n")},
		{9, strconv.Quote("1\n1\n2\n3\n5\n8\n13\n21\n34\n")},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		out = &buf

		fib(test.input)

		expected := test.expected
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Incorrect output from fib()")
			t.Errorf("expected: %v\nactual: %v", expected, actual)
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output from main()")
		t.Errorf("expected: %v\nactual: %v", expected, actual)
	}
}
