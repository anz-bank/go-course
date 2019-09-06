package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNotModifiedInput(t *testing.T) {
	input := []int{64, 11, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	original := make([]int, len(input))
	copy(original, input)

	bubble(input)

	if !reflect.DeepEqual(original, input) {
		t.Errorf("The input slice was altered %v does not match %v", original, input)
	}
}
func TestBubble(t *testing.T) {
	input := []int{64, 11, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	expected := []int{1, 1, 2, 3, 5, 8, 11, 13, 21, 34, 64}

	sortedResult := bubble(input)

	if !reflect.DeepEqual(sortedResult, expected) {
		t.Errorf("Returned value %v does not match %v", sortedResult, expected)
	}
}
func TestNoInput(t *testing.T) {
	input := []int{}

	sortedResult := bubble(input)

	if len(sortedResult) != 0 {
		t.Errorf("Returned value %v but expected empty slice", sortedResult)
	}
}
func TestNoSort(t *testing.T) {
	input := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

	sortedResult := bubble(input)

	if !reflect.DeepEqual(sortedResult, expected) {
		t.Errorf("Returned value %v does not match %v", sortedResult, expected)
	}
}
func TestMain(t *testing.T) {
	expected := "[1 2 3 5]\n"
	var buf bytes.Buffer
	out = &buf

	main()

	result := buf.String()

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
