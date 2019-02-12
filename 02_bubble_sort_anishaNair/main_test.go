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
	expected := "[1 2 3 10 13 14 15 50]\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleOutput(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubble([]int{3, 2, 1, 15, 13, 14, 50, 10})

	// Then
	expected := []int{1, 2, 3, 10, 13, 14, 15, 50}
	actual := bubble([]int{3, 2, 1, 15, 13, 14, 50, 10})
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestBubbleOutputWithNegativeNumbers(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubble([]int{-3, 2, -1, 15, 13, 14, 50, 10})

	// Then
	expected := []int{-3, -1, 2, 10, 13, 14, 15, 50}
	actual := bubble([]int{-3, 2, -1, 15, 13, 14, 50, 10})
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleOutputWithZeros(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubble([]int{0, 0})

	// Then
	expected := []int{0, 0}
	actual := bubble([]int{0, 0})
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleOutputWithEmptyArray(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubble([]int{})

	// Then
	expected := []int{}
	actual := bubble([]int{})
	r.Equalf(expected, actual, "Unexpected output in main()")
}
