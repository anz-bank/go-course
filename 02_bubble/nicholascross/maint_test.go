package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	expectSorted([]int{1,2,3,4,5}, []int{1,2,3,4,5}, t)
	expectSorted([]int{5,4,3,2,1}, []int{1,2,3,4,5}, t)
	expectSorted([]int{1,2,4,3,5}, []int{1,2,3,4,5}, t)
	expectSorted([]int{4,3,1,5,2}, []int{1,2,3,4,5}, t)
	expectSorted([]int{5,1,2,4,3}, []int{1,2,3,4,5}, t)
}

func expectSorted(original []int, expected []int, t *testing.T) {
	actual := bubble(original)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Sort slice %v does not match expected slice %v", actual, expected)
	} else {
		t.Logf("Sorted %v -> %v", original, actual)
	}
}
