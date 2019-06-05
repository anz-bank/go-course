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
func TestBubble(t *testing.T) {

	input := []int{3, 2, 1, 5}

	got := bubbleSort(input)
	want := []int{1, 2, 3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given %v", got, want, input)
	}
}

func TestInsertion(t *testing.T) {

	input := []int{3, 2, 1, 5}

	got := insertionSort(input)
	want := []int{1, 2, 3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given %v", got, want, input)
	}
}
