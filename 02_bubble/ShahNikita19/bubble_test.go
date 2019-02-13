package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `[1 2 3 5]`
	actual := strings.TrimSuffix(buf.String(), "\n")
	r.Equalf(expected, actual, "Unexpected output in main() ")
}

func TestBubbleSortWithEmptyArray(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubbleSort([]int{})

	// Then
	expected := []int{}
	actual := bubbleSort([]int{})
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleSortWithNegativeNumbers(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubbleSort([]int{-3, 2, -1, 15, 13, 14, 50, 10})

	// Then
	expected := []int{-3, -1, 2, 10, 13, 14, 15, 50}
	actual := bubbleSort([]int{-3, 2, -1, 15, 13, 14, 50, 10})
	r.Equalf(expected, actual, "Unexpected output in bubble()")
}
func TestBubbleSortWithZeros(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	bubbleSort([]int{0, 0})

	// Then
	expected := []int{0, 0}
	actual := bubbleSort([]int{0, 0})
	r.Equalf(expected, actual, "Unexpected output in bubble()")
}

func TestBubbleSortWithOriginalValue(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	originalList := []int{3, 2, 1, 5}
	bubbleSort(originalList)

	// Then
	expected := []int{3, 2, 1, 5}
	r.Equalf(expected, originalList, "Unexpected output in Insertion()")
}
