package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//when
	main()

	//Then
	expected := "[1 2 3 5]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output")
}

var testCases = []struct {
	in  []int
	out []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{-3, -2, -1, -5}, []int{-5, -3, -2, -1}},
	{[]int{5}, []int{5}},
	{[]int{-10, 1, 5, -16, 18}, []int{-16, -10, 1, 5, 18}},
	{[]int{1, 1, -6, 8}, []int{-6, 1, 1, 8}},
}

func TestBubbleSort(t *testing.T) {
	for _, arr := range testCases {
		// When
		actual := bubble(arr.in)

		// Then
		require.Equal(t, arr.out, actual, "Unexpected output")
	}
}

func TestHeapSort(t *testing.T) {
	for _, arr := range testCases {
		// When
		actual := heapSort(arr.in)

		// Then
		require.Equal(t, arr.out, actual, "Unexpected output")
	}
}
