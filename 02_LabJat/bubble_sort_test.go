package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInRageOutputBubbleSort(t *testing.T) {
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

	actual := 0
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
}
func TestSingleElementArrayBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	outPut := bubbleSort([]int{1})

	// Then

	actual := 1
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
}

func TestSingleElementArrayInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	outPut := insertionSort([]int{1})

	// Then

	actual := 1
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
}

func TestEmptyArrayInsertionSort(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	outPut := insertionSort([]int{})

	// Then

	actual := 0
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
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
	actual1 := 1
	actual2 := 5
	r.Equalf(outPut1, actual1, "Unexpected output in main()")
	r.Equalf(outPut2, actual2, "Unexpected output in main()")
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
