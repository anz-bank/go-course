package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortEdgeCases(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{}, bubble([]int{}), "empty slice as input")
	assert.Equal([]int{4}, bubble([]int{4}), "one integer as input")
	assert.Equal([]int{-8, 3}, bubble([]int{3, -8}), "two numbers as input")
}

func TestSortInputOptions(t *testing.T) {
	assert := assert.New(t)
	longInput := []int{999, 888, 777, 666, 555, 444, 333, 222, 111, 99, 88, 77, 66, 55, 44, 33, 22, 11, 1}
	longOutput := []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99, 111, 222, 333, 444, 555, 666, 777, 888, 999}

	assert.Equal([]int{1, 11, 111, 222, 333}, bubble([]int{1, 11, 111, 222, 333}), "input is already sorted")
	assert.Equal([]int{11, 22, 33, 44}, bubble([]int{44, 33, 22, 11}), "input is reverse sorted")
	assert.Equal([]int{-88, -3, 0, 3, 88}, bubble([]int{88, -3, -88, 0, 3}), "positive and negative numbers")
	assert.Equal([]int{-5, -5, -5, -2, -2, 3}, bubble([]int{-2, 3, -2, -5, -5, -5}), "duplicate numbers in input")
	assert.Equal(longOutput, bubble(longInput), "sort 19 numbers")
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[1 2 3 5]\n", buf.String())
}
