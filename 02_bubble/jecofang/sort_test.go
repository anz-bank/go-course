package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
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

var testData = []struct {
	in  []int
	out []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, -1}, []int{-1, 1, 2, 3, 5}},
	{[]int{3, 2, 1, 5, -1, 3}, []int{-1, 1, 2, 3, 3, 5}},
	{[]int{}, []int{}},
	{[]int{10}, []int{10}},
	{[]int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

func TestBubbleOutput(t *testing.T) {
	for _, d := range testData {
		// When
		actual := bubble(d.in)

		// Then
		require.Equal(t, d.out, actual, "Unexpected output")
	}
}

func TestInsertionOutput(t *testing.T) {
	for _, d := range testData {
		// When
		actual := insertion(d.in)

		// Then
		require.Equal(t, d.out, actual, "Unexpected output")
	}
}

func TestQuickOutput(t *testing.T) {
	for _, d := range testData {
		//When
		actual := quick(d.in)

		//Then
		require.Equal(t, d.out, actual, "Unexpected output")
	}
}
