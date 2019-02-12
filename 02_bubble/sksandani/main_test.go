package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	i := bubble([]int{3, 2, 1, 5}) // Then
	expected := []int{1, 2, 3, 5}
	actual := i
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	i := insertion([]int{3, 2, 1, 5}) // Then
	expected := []int{1, 2, 3, 5}
	actual := i
	r.Equalf(expected, actual, "Unexpected output in main()")
}
