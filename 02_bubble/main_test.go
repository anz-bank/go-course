package main

import (
	"bytes"
	"strconv"
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
	expected := strconv.Quote("After sorting =  [1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
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
}

func TestBubbleSourt(t *testing.T) {
	r := require.New(t)
	for _, tt := range testData {
		bubbleSort(tt.in)
		r.Equalf(tt.out, tt.in, "Test case fails")
	}
}
