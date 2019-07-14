package main

import (
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestBubbleWithSmallSlice(t *testing.T) {
	a := assert.New(t)

	sorted := bubble([]int{2, 1})
	expected := []int{1, 2}

	a.Equal(expected, sorted)
}

func TestBubbleWithRepeatedNumbers(t *testing.T) {
	a := assert.New(t)

	sorted := bubble([]int{3, 1, 2, 3, 2, 1})
	expected := []int{1, 1, 2, 2, 3, 3}

	a.Equal(expected, sorted)
}

func TestBubbleWithLargeSlice(t *testing.T) {
	a := assert.New(t)

	expected := make([]int, 100)
	for i := range expected {
		expected[i] = rand.Intn(1000)
	}

	sorted := bubble(expected)
	sort.Ints(expected)

	a.Equal(expected, sorted)
}

func TestBubbleWithSmallSliceIncludingNegativeInts(t *testing.T) {
	a := assert.New(t)

	sorted := bubble([]int{1, -2, -3})
	expected := []int{-3, -2, 1}

	a.Equal(expected, sorted)
}
