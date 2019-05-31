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
	"Empty array":                 {input: []int{}, expected: []int{}},
	"Single element array":        {input: []int{101}, expected: []int{101}},
	"Best case (sorted)":          {input: []int{1, 2, 3, 4, 5, 6, 7}, expected: []int{1, 2, 3, 4, 5, 6, 7}},
	"Worst Case (reverse sorted)": {input: []int{7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7}},
	"Average Case (random order)": {input: []int{59, 62, 48, 53, 100, 86, 10, 68, 50, 9, 2},
		expected: []int{2, 9, 10, 48, 50, 53, 59, 62, 68, 86, 100}},
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

func TestBubbleSort(t *testing.T) {
	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := bubble(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", testData.expected, actual)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := insertionSort(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", testData.expected, actual)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := MergeSort(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", testData.expected, actual)
			}
		})
	}
}
