package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in  []int
	out []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{-3, -2, -1, -5}, []int{-5, -3, -2, -1}},
	{[]int{}, []int{}},
	{[]int{5}, []int{5}},
	{[]int{5, 3, 2, 1}, []int{1, 2, 3, 5}},
	{[]int{5, 3, 5, 1}, []int{1, 3, 5, 5}},
}

func TestBubbleOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		out := bubble(tt.in)
		r.ElementsMatch(tt.out, out)
	}
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[1 2 3 5]\n"
	r.Equalf(expected, buf.String(), "Unexpected output in main()")
}

func TestBubbleDoesntModifyInput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	arr := []int{-3, -2, -1, -5}
	actual := bubble(arr)

	// Then
	r.Equal([]int{-3, -2, -1, -5}, arr)
	r.Equal([]int{-5, -3, -2, -1}, actual)
}
