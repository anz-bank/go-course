package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[1 2 3 5]"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}

func TestBubbleAndInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Empty slice", input: []int{}, expected: []int{}},
		{name: "Reverse sorted slice", input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "Already sorted slice", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "Slice with one number different to rest", input: []int{1, 1, 1, 2, 1},
			expected: []int{1, 1, 1, 1, 2}},
		{name: "Slice with negative number", input: []int{1, -1}, expected: []int{-1, 1}},
	}

	for _, test := range tests {
		actual := bubble(test.input)
		testResult := equal(actual, test.expected)
		if testResult == false {
			t.Fatalf("Bubble Sort Test Failed! Expected: %v, Got: %v", test.expected, actual)
		} else {
			t.Logf("Bubble Sort Test Passed: %v", test.name)
		}
	}

	for _, test := range tests {
		actual := insertionSort(test.input)
		testResult := equal(actual, test.expected)
		if testResult == false {
			t.Fatalf("Insertion Sort Test Failed! Expected: %v, Got: %v", test.expected, actual)
		} else {
			t.Logf("Insertion Sort Test Passed: %v", test.name)
		}
	}
}

func TestIntSliceEqual(t *testing.T) {
	tests := []struct {
		name     string
		sliceA   []int
		sliceB   []int
		expected bool
	}{
		{name: "Two slices with different length", sliceA: []int{1, 0}, sliceB: []int{1}, expected: false},
		{name: "Two slices are different", sliceA: []int{1, 0}, sliceB: []int{1, 1}, expected: false},
		{name: "Two slices are equal", sliceA: []int{1, 1}, sliceB: []int{1, 1}, expected: true},
	}

	for _, test := range tests {
		t.Logf("Running test: %v", test.name)
		actual := equal(test.sliceA, test.sliceB)
		if actual != test.expected {
			t.Fatalf("Test failed! Expected: %t, Got: %t", test.expected, actual)
		}
	}
}
