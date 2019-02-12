package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

//To test the bubbleFunction with positive numbers
func TestBubble_PositiveNumbers(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	b := []int{5, 2, 3, 4, 1}
	actual := bubble(b)
	expected := []int{1, 2, 3, 4, 5}
	r.Equalf(expected, actual, "Unexpected output in bubble")
}

//To test the bubbleFunction negative numbers
func TestBubble_NegativeNumbers(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	b := []int{-5, 2, 3, 0, 1, 4}
	actual := bubble(b)
	expected := []int{-5, 0, 1, 2, 3, 4}
	r.Equalf(expected, actual, "Unexpected output in bubble")
}

//To test the copying
func TestBubble_DoesntModifyInput(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	b := []int{-5, 2, 3, 0, 1, 4}
	actual := bubble(b)

	// Then
	r.Equal([]int{-5, 0, 1, 2, 3, 4}, actual)
	r.Equal([]int{-5, 2, 3, 0, 1, 4}, b)

}

//To test the main output
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
