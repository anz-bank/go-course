package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestBubbleOutput(t *testing.T) {

	testCases := map[string]struct {
		input []int
		want  []int
	}{
		"normal": {input: []int{3, 2, 1, 5}, want: []int{1, 2, 3, 5}},
		"blank":  {input: []int{}, want: []int{}},
		"nil":    {input: nil, want: nil},
		"duplicated":    {input: []int{1, 0, 0, 1}, want: []int{0, 0, 1, 1}},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		test := test

		t.Run(name, func(t *testing.T) {
			actual := bubble(test.input)

			if !reflect.DeepEqual(test.want, actual) {
				t.Errorf("running : %v, expected %v, got %v", test.input, test.want, actual)
			}
		})
	}
}

func TestInsertSortOutput(t *testing.T) {

	testCases := map[string]struct {
		input []int
		want  []int
	}{
		"normal": {input: []int{3, 2, 1, 5}, want: []int{1, 2, 3, 5}},
		"blank":  {input: []int{}, want: []int{}},
		"nil":    {input: nil, want: nil},
		"duplicated":    {input: []int{1, 0, 0, 1}, want: []int{0, 0, 1, 1}},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		test := test

		t.Run(name, func(t *testing.T) {
			actual := insertSort(test.input)

			if !reflect.DeepEqual(test.want, actual) {
				t.Errorf("running : %v, expected %v, got %v", test.input, test.want, actual)
			}
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("running main, expected %v, got %v", expected, actual)
	}
}
