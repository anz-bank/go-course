package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[1 2 3 5]\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var testData = []struct {
	in  []int
	out []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, -1}, []int{-1, 1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, -1, 3}, []int{-1, 1, 2, 3, 3, 5}},
	{[]int{3, 2, 1, 5, -1, 3, 0}, []int{-1, 0, 1, 2, 3, 3, 5}},
	{[]int{}, []int{}},
	{[]int{10}, []int{10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestBubbleSort(t *testing.T) {
	r := require.New(t)
	for _, tt := range testData {
		result := bubbleSort(tt.in)
		r.Equalf(tt.out, result, "Test case fails")
	}
}

func TestBubbleSortInput(t *testing.T) {
	r := require.New(t)
	for _, tt := range testData {
		input := make([]int, len(tt.in))
		copy(input, tt.in)
		bubbleSort(tt.in)
		r.Equalf(input, tt.in, "Test case fails")
	}
}
