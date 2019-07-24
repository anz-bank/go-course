package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestBubbleSortFunc(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected []int
	}

	tests := []test{
		{name: "already sorted", input: []int{2, 4, 5, 6, 8, 10}, expected: []int{2, 4, 5, 6, 8, 10}},
		{name: "descending order", input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "same number", input: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
		{name: "empty array", input: []int{}, expected: []int{}},
		{name: "average case", input: []int{59, 62, 48, 53, 100, 86, 10, 68, 50, 9, 2},
			expected: []int{2, 9, 10, 48, 50, 53, 59, 62, 68, 86, 100}},
	}

	for _, testCase := range tests {
		actual := bubbleSort(testCase.input)
		fmt.Println(testCase.name, "expected: ", testCase.expected)
		fmt.Println(testCase.name, "actual: ", actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
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
		t.Errorf("TestMainOutput: Unexpected output in main()")
	}
}
