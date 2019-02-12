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
	outWriter = &buf

	// When
	main()

	// Then
	expected := "[1 2 3 5]"
	actual := buf.String()

	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := insertion([]int{3, 2, 1, 5})

	// Then
	expected := []int{1, 2, 3, 5}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}

func TestInsertionSortWithNegativeInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := insertion([]int{3, 2, 1, -4, -9, 5, 0})

	// Then
	expected := []int{-9, -4, 0, 1, 2, 3, 5}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}

func TestBubbleSortWithNegativeInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := bubble([]int{3, 2, 1, -4, -9, 5, 0})

	// Then
	expected := []int{-9, -4, 0, 1, 2, 3, 5}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}

func TestBubbleSortRepeatedInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := bubble([]int{3, 2, 1, -4, -9, 5, 0, 3, 2})

	// Then
	expected := []int{-9, -4, 0, 1, 2, 2, 3, 3, 5}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}
func TestInsertionSortRepeatedInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := insertion([]int{3, 2, 1, -4, -9, 5, 0, 3, 2})

	// Then
	expected := []int{-9, -4, 0, 1, 2, 2, 3, 3, 5}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}

func TestBubbleSortWithSingleInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := bubble([]int{3})

	// Then
	expected := []int{3}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}

func TestInsertionSortWithSingleInt(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	sortedList := insertion([]int{3})

	// Then
	expected := []int{3}
	actual := sortedList
	r.Equalf(expected, actual, "Unexpected output in Insertion()")
}
