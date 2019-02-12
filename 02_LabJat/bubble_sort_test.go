package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInRageOutput(t *testing.T) {
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
func TestEmptyArray(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	outPut := bubbleSort([]int{})

	// Then

	actual := 0
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
}
func TestSingleElementArray(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	outPut := bubbleSort([]int{1})

	// Then

	actual := 1
	r.Equalf(len(outPut), actual, "Unexpected output in main()")
}
