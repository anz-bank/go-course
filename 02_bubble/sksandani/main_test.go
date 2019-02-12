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

func TestBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	i := bubble([]int{3, 2, 1, 5}) // Then
	expected := []int{1, 2, 3, 5}
	actual := i
	r.Equalf(expected, actual, "Unexpected output in Bubble()")
}

func TestInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	i := insertion([]int{3, 2, 1, 5}) // Then
	expected := []int{1, 2, 3, 5}
	actual := i
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}
