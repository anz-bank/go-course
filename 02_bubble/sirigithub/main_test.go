package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

var tests = map[string]struct {
	input    []int
	expected []int
}{
	"Empty":                   {input: []int{}, expected: []int{}},
	"Single element":          {input: []int{1}, expected: []int{1}},
	"Already Sorted":          {input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
	"Sorted in reverse order": {input: []int{10, 9, 8, 7, 6}, expected: []int{6, 7, 8, 9, 10}},
	"Negative elements":       {input: []int{-10, 9, 8, 7, 6}, expected: []int{-10, 6, 7, 8, 9}},
	"Random order":            {input: []int{3, 2, 1, 5}, expected: []int{1, 2, 3, 5}},
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
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := insertionSort(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("actual: %d but expected: %d ", test.expected, actual)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := bubbleSort(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("actual: %d but expected: %d ", test.expected, actual)
			}
		})
	}
}
