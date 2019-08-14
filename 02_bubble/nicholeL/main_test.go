package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		excepted []int
	}{
		"one":   {[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
		"two":   {[]int{3, 2, 1, 7, 9, 5}, []int{1, 2, 3, 5, 7, 9}},
		"three": {[]int{3, 2, 1, 7, 9, 5, 10}, []int{1, 2, 3, 5, 7, 9, 10}},
		"four":  {[]int{3, 2, 1, 7, 9, 5, 10, 12}, []int{1, 2, 3, 5, 7, 9, 10, 12}},
		"five":  {[]int{}, []int{}},
	}

	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := bubbleSort(test.input)
			if !reflect.DeepEqual(test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})

	}

}

func TestInsertSort(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		excepted []int
	}{
		"one":   {[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
		"two":   {[]int{3, 2, 1, 7, 9, 5}, []int{1, 2, 3, 5, 7, 9}},
		"three": {[]int{3, 2, 1, 7, 9, 5, 10}, []int{1, 2, 3, 5, 7, 9, 10}},
		"four":  {[]int{3, 2, 1, 7, 9, 5, 10, 12}, []int{1, 2, 3, 5, 7, 9, 10, 12}},
		"five":  {[]int{}, []int{}},
	}

	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := insertSort(test.input)
			if !reflect.DeepEqual(test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})

	}

}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	excepted := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
	if excepted != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
	}
}
