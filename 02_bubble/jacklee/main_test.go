package main

import (
	"bytes"
	"strconv"
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
	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
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
	// Given
	r := require.New(t)
	for _, arr := range testCases {
		in := make([]int, len(arr.in))
		copy(in, arr.in)

		// When
		actual := bubble(in)

		// Then
		r.Equalf(arr.out, actual, "Unexpected output")
	}
}

func TestHeapSort(t *testing.T) {
	// Given
	r := require.New(t)
	for _, arr := range testCases {
		in := make([]int, len(arr.in))
		copy(in, arr.in)

		// When
		actual := heapSort(in)

		// Then
		r.Equalf(arr.out, actual, "Unexpected output")
	}
}
