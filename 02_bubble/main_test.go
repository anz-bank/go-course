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
