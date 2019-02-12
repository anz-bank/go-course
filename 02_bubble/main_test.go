package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := bubble([]int{3, 2, 1, 5})

	// Then
	expected := []int([]int{1, 2, 3, 5})
	r.Equalf(expected, s, "Unexpected output in main()")
}

func TestBubbleOutputForNegativeNumbers(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := bubble([]int{-3, -2, -1, -5})

	// Then
	expected := []int{-5, -3, -2, -1}
	r.Equalf(expected, s, "Unexpected output in main()")
}
