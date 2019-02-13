package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected := "[1 2 3 5]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestBubbleSortNegativeValues(t *testing.T) {
	//Given
	assert := assert.New(t)
	//when
	input := []int{-3, -2, -1, -5}
	sortedArray := bubble(input)
	//Then
	assert.Equal([]int{-5, -3, -2, -1}, sortedArray)
}

func TestSliceCopyFailure(t *testing.T) {
	//Given
	assert := assert.New(t)
	//when
	input := []int{1, 2, 3, 5}
	sortedArray := bubble(input)
	//Then
	assert.NotEqual(assert, input, sortedArray)
}
