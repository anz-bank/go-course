package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainInputOutput(t *testing.T) {
	// Given
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

func TestBubbleSortOutput(t *testing.T) {
	// Given
	r := require.New(t)

	// Then
	expected := []int{1, 2, 3, 5}
	actual := bubble([]int{3, 2, 1, 5})
	r.EqualValues(expected, actual, "Unexpected output in bubble()")
}

func TestInsertionSortOutput(t *testing.T) {
	// Given
	r := require.New(t)

	// Then
	expected := []int{1, 2, 3, 5}
	actual := insertion([]int{3, 2, 1, 5})
	r.EqualValues(expected, actual, "Unexpected output in insertion()")
}

var tests = []struct {
	in          []int
	expectedOut []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2, 3, 5}, []int{1, 2, 3, 5}},
	{[]int{5, 3, 2, 1}, []int{1, 2, 3, 5}},
	{[]int{5, 2, 5, 1}, []int{1, 2, 5, 5}}}

func TestBubbleSortEdgeCases(t *testing.T) {
	r := require.New(t)

	for _, tt := range tests {
		expectedIn := make([]int, len(tt.in))
		copy(expectedIn, tt.in)
		out := bubble(tt.in)
		r.Equal(expectedIn, tt.in) // ensure tt.in has not been modified
		r.Equal(tt.expectedOut, out)
	}
}

func TestInsertionSortEdgeCases(t *testing.T) {
	r := require.New(t)

	for _, tt := range tests {
		expectedIn := make([]int, len(tt.in))
		copy(expectedIn, tt.in)
		out := insertion(tt.in)
		r.Equal(expectedIn, tt.in) // ensure tt.in has not been modified
		r.Equal(tt.expectedOut, out)
	}
}
