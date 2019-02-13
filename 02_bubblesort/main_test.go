package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBubbleSortOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	t.Log(actual)
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleSortOutputWithNegValues(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	s := bubble([]int{-3, 2, -1, 5})

	// Then
	assert.Equal([]int{-3, -1, 2, 5}, s)
}

func TestBubbleSortOutputWithEmptySlices(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	s := bubble([]int{})

	// Then
	assert.Equal(0, len(s))
}

func TestBubbleSortOutputWithDuplicates(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	s := bubble([]int{2, 2, 1, 5})

	// Then
	assert.Equal([]int{1, 2, 2, 5}, s)
}

func TestBubbleSortCopyFailure(t *testing.T) {
	// Given
	assert := assert.New(t)
	// When
	input := []int{2, 2, 1, 5}
	sortedArray := bubble(input)
	// Then
	assert.NotEqual(input, sortedArray)	
}
