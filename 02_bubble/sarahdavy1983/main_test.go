package main

import (
	"reflect"
	"testing"
)

func TestMainOut(t *testing.T) {

	main()

	expected := []int{1, 2, 3, 5}
	actual := []int{1, 2, 3, 5}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()")
	}
}

var tests = []struct {
	input []int
	want  []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
	{[]int{}, []int{}},
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
}

func TestBubble(t *testing.T) {
	for _, tc := range tests {
		got := (bubbleSort(tc.input))
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("got %v want %v given %v", got, tc.want, tc.input)
		}
	}
}

func TestInsertion(t *testing.T) {
	for _, tc := range tests {
		got := insertionSort(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v want %v given %v", got, tc.want, tc.input)
		}
	}
}
