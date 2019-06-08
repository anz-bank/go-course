package main

import (
	"reflect"
	"testing"
)

var bubbleSorttests = []struct {
	input []int
	want  []int
}{
	{input: []int{3, 2, 1, 5}, want: []int{1, 2, 3, 5}},
	{input: []int{}, want: []int{}},
	{input: nil, want: nil},
	{input: []int{1, 0, 0, 1}, want: []int{0, 0, 1, 1}},
	{input: []int{1}, want: []int{1}},
}

func TestMainOutput(t *testing.T) {
	for _, test := range bubbleSorttests {
		actual := bubbleSort(test.input)
		if !reflect.DeepEqual(test.want, actual) {
			t.Errorf("running : %v, expected %v, got %v", test.input, test.want, actual)
		}
	}
}
