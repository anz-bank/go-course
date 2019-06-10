package main

import (
  "bytes"
	"strconv"
	"testing"
)

//test cases to write : empty array, single element, sorted reverse,some negative elements
var tests = []struct {
	input    []int
	expected []int
}{
	{input: []int{}, expected: []int{}},
	{input: []int{1}, expected: []int{1}},
	{input: []int{1,2,3,4,5}, expected: []int{1,2,3,4,5}},
	{input: []int{10,9,8,7,6}, expected: []int{6,7,8,9,10}},
	{input: []int{-10,9,8,7,6}, expected: []int{-10,6,7,8,9}},
	{input: []int{3, 2, 1, 5}, expected: []int{1,2,3,5}},
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestInsertionSort(t *testing.T) {

	for _, test := range tests {
		actual := insertionSort(test.input)
		if !(Equal(actual, test.expected)) {
			t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", test.expected, actual)
		}
	}
}

func TestBubbleSort(t *testing.T) {

	for _, test := range tests {
		actual := bubbleSort(test.input)
		if !(Equal(actual, test.expected)) {
			t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", test.expected, actual)
		}
	}
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
