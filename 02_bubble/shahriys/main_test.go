package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {

	type test struct {
		input    []int
		expected []int
		actual   []int
	}

	tests := []test{
		{input: []int{3, 2, 1, 5}, expected: []int{1, 2, 3, 5}, actual: bubble([]int{3, 2, 1, 5})},
		{input: []int{}, expected: []int{}, actual: bubble([]int{})},
		{input: []int{3, 2, 1, 5}, expected: []int{1, 2, 3, 5}, actual: insertion([]int{3, 2, 1, 5})},
		{input: []int{}, expected: []int{}, actual: insertion([]int{})},
	}

	for _, tc := range tests {

		if !reflect.DeepEqual(tc.expected, tc.actual) {
			t.Fatalf(" input:%v, expected: %v, got: %v", tc.input, tc.expected, tc.actual)
		}
	}

}

func TestBubbleSortMain(t *testing.T) {
	var bufbub bytes.Buffer
	outbub = &bufbub

	main()
	expected := `[1 2 3 5][1 2 3 5]`
	actual := bufbub.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()" + expected + " " + actual)
	}

}
