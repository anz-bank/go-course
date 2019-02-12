package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var sortingTestsData = []struct {
	n        []int // input
	expected []int // expected result
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, 8}, []int{1, 2, 3, 5, 8}},
}

func TestBubbleSortMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[1 2 3 5]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleSortOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range sortingTestsData {
		actual := bubble(tt.n)
		r.ElementsMatchf(tt.expected, actual, "Unexpected output in main()")
	}
}

func TestInsertionSortOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range sortingTestsData {
		actual := insertionsort(tt.n)
		if reflect.DeepEqual(actual, tt.expected) {
			r.ElementsMatchf(tt.expected, actual, "Unexpected output in main()")
		}
	}
}
