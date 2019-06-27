package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleEdgeCases(t *testing.T) {
	assert.Equal(t, []int{}, bubble([]int{}))
	assert.Equal(t, []int{1}, bubble([]int{1}))
	assert.Equal(t, []int{1, 2}, bubble([]int{2, 1}))
}

func TestBubbleBestCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, bubble([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestBubbleAverageCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 6, 9, 10, 12}, bubble([]int{12, 2, 1, 6, 9, 3, 10}))
}

func TestBubbleWorstCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, bubble([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestBubbleDuplicates(t *testing.T) {
	assert.Equal(t, []int{2, 3, 3, 6, 6, 8, 9}, bubble([]int{3, 2, 3, 6, 8, 6, 9}))
}

func TestBubbleDoesNotTouchInput(t *testing.T) {
	ns := []int{3, 2, 1}
	bubble(ns)
	assert.Equal(t, []int{3, 2, 1}, ns)
}

func TestInsertionEdgeCases(t *testing.T) {
	assert.Equal(t, []int{}, insertion([]int{}))
	assert.Equal(t, []int{1}, insertion([]int{1}))
	assert.Equal(t, []int{1, 2}, insertion([]int{2, 1}))
}

func TestInsertionBestCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, insertion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestInsertionAverageCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 6, 9, 10, 12}, insertion([]int{12, 2, 1, 6, 9, 3, 10}))
}

func TestInsertionWorstCase(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, insertion([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestInsertionDuplicates(t *testing.T) {
	assert.Equal(t, []int{2, 3, 3, 6, 6, 8, 9}, insertion([]int{3, 2, 3, 6, 8, 6, 9}))
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[1 2 3]\n[1 2 3]\n", buf.String())
}
