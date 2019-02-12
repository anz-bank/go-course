package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleOutput(t *testing.T) {
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

var tests = []struct {
	in  []int
	out []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, -1, -90, -123, 4387, 29384}, []int{-123, -90, -1, 1, 2, 3, 5, 4387, 29384}},
	{[]int{8, 5, 3, 1, 9, 6, 0, 7, 4, 2, 5}, []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}}

func TestBubbleSortWithValues(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		out := bubble(tt.in)
		r.ElementsMatch(tt.out, out)
	}
}

func TestInsertionSortWithValues(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		fmt.Println("insertion testing", tt.in, tt.out)
		out := insertionSort(tt.in)
		r.ElementsMatch(tt.out, out)
	}
}
