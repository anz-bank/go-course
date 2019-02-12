package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRageOutputBubbleSort(t *testing.T) {
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
func TestEmptyArrayBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)
	// When
	outPut := bubbleSort([]int{})
	// Then
	expected := 0
	r.Equalf(expected, len(outPut), "Unexpected output in main()")
}
func TestSingleElementArrayBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)
	// When
	outPut := bubbleSort([]int{1})
	// Then
	expected := 1
	r.Equalf(expected, len(outPut), "Unexpected output in main()")
}

func TestSingleElementArrayInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)
	// When
	outPut := insertionSort([]int{1})
	// Then
	expected := 1
	r.Equalf(expected, len(outPut), "Unexpected output in main()")
}

func TestEmptyArrayInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)
	// When
	outPut := insertionSort([]int{})
	// Then
	expected := 0
	r.Equalf(expected, len(outPut), "Unexpected output in main()")
}

func TestRangeValuesArrayInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	outPut := insertionSort([]int{3, 2, 1, 5})
	// Then
	outPut1 := outPut[0]
	outPut2 := outPut[3]
	expected1 := 1
	expected2 := 5
	r.Equalf(expected1, outPut1, "Unexpected output in main()")
	r.Equalf(expected2, outPut2, "Unexpected output in main()")
}

func TestInRageOutputInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	outPut := insertionSort([]int{3, 2, 1, 5})
	fmt.Fprintln(out, outPut)
	// Then
	expected := "[1 2 3 5]\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestRageOutputInsertionSortElementsMatch(t *testing.T) {
	// Given
	var buf bytes.Buffer
	out = &buf
	// When
	outPut := insertionSort([]int{7, 4, 3, 2, 1, 5})
	fmt.Fprintln(out, outPut)
	// Then
	expected := []int{7, 4, 3, 2, 1, 5}
	assert.ElementsMatch(t, expected, outPut)
}
