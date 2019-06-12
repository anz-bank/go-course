package main

import (
	"bytes"
	"strconv"
	"testing"
)

// table driven tests
var tests = []struct {
	input    []int
	expected []int
}{
	{
		[]int{3, 2, 1, 5},
		[]int{1, 2, 3, 5},
	},
	{
		[]int{},
		[]int{},
	},
	{
		[]int{1},
		[]int{1},
	},
	{
		[]int{1, 2},
		[]int{1, 2},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, test := range tests {
		testSort(t, bubble(test.input), test.expected)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range tests {
		testSort(t, insertion(test.input), test.expected)
	}
}

func testSort(t *testing.T, actual []int, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("actual and expected slices have different lengths")
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Unexpected result")
			t.Errorf("\nActual: %v\nExpected: %v", actual, expected)
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
		t.Errorf("\nActual: %q\nExpected: %q", actual, expected)
	}
}
