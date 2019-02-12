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
	expected := `[1 2 3 5][1 2 3 5]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleSortOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	bubble([]int{3, 2, 1, 5})

	// Then
	expected := `[1 2 3 5]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestInsertionSortOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	insertion([]int{3, 2, 1, 5})

	// Then
	expected := `[1 2 3 5]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
