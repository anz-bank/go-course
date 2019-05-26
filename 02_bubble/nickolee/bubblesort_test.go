package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestBubbleSortFunc(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	// Test empty slice
	test1 := bubbleSort([]int{})
	test1Expected := []int{}

	// Test random slice
	test2 := bubbleSort([]int{3, 2, 5, 1, 7, 2, 9, 1, 0})
	test2Expected := []int{0, 1, 1, 2, 2, 3, 5, 7, 9}

	// Check whether expected results equal actual results
	if !reflect.DeepEqual(test1, test1Expected) || !reflect.DeepEqual(test2, test2Expected) {
		t.Errorf("TestBubbleSortFunc: Unexpected output in main()")
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
