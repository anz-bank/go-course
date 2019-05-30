package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	expectSorted([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, bubble, t)
	expectSorted([]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}, bubble, t)
	expectSorted([]int{1, 2, 4, 3, 5}, []int{1, 2, 3, 4, 5}, bubble, t)
	expectSorted([]int{4, 3, 1, 5, 2}, []int{1, 2, 3, 4, 5}, bubble, t)
	expectSorted([]int{5, 1, 2, 4, 3}, []int{1, 2, 3, 4, 5}, bubble, t)
}

func TestInsertionSort(t *testing.T) {
	expectSorted([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, insertion, t)
	expectSorted([]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}, insertion, t)
	expectSorted([]int{1, 2, 4, 3, 5}, []int{1, 2, 3, 4, 5}, insertion, t)
	expectSorted([]int{4, 3, 1, 5, 2}, []int{1, 2, 3, 4, 5}, insertion, t)
	expectSorted([]int{5, 1, 2, 4, 3}, []int{1, 2, 3, 4, 5}, insertion, t)
}

func expectSorted(original []int, expected []int, sort func(s []int) []int, t *testing.T) {
	actual := sort(original)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Sort slice %v does not match expected slice %v", actual, expected)
	} else {
		t.Logf("Sorted %v -> %v", original, actual)
	}
}
