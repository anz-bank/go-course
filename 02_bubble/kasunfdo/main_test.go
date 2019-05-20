package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

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

func TestInsertionSort(t *testing.T) {
	input := []int{3, 2, 1, 5}

	expected := []int{1, 2, 3, 5}
	actual := insertionSort(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestMergeSort(t *testing.T) {
	input := []int{3, 2, 1, 5}

	expected := []int{1, 2, 3, 5}
	actual := MergeSort(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
