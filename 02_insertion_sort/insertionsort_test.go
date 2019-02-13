package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertionSortOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()
	// Then
	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestInsertionSortOutputWithNegValues(t *testing.T) {
	// Given
	assert := assert.New(t)
	// When
	s := insertion([]int{-3, 2, -1, 5})
	// Then
	assert.Equal([]int{-3, -1, 2, 5}, s)
}

func TestBubbleSortCopyFail(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	input := []int{3, 2, 1, 5}
	sortedArray := insertion(input)
	
	// Then
	assert.NotEqual(input, sortedArray)	
}
