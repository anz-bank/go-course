package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "Sorting failed.")
}

var testCases = map[string]struct {
	input []int
	want  []int
}{
	"main":                {input: []int{3, 1, 2, 5}, want: []int{1, 2, 3, 5}},
	"ascending sequence":  {input: []int{1, 2, 4, 3, 5}, want: []int{1, 2, 3, 4, 5}},
	"descending sequence": {input: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	"duplicates":          {input: []int{4, 4, 1, 100, 1}, want: []int{1, 1, 4, 4, 100}},
	"negatives":           {input: []int{-5, -1, -2, 4, 3, 19, 0, 1, 4}, want: []int{-5, -2, -1, 0, 1, 3, 4, 4, 19}},
	"empty":               {input: []int{}, want: []int{}},
	"single":              {input: []int{1}, want: []int{1}},
}

func TestBubbleSort(t *testing.T) {
	for _, test := range testCases {
		input := test.input
		expected := test.want
		sorted := bubble(input)
		assert.Equalf(t, expected, sorted, "Bubble sort failed")
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range testCases {
		input := test.input
		expected := test.want
		sorted := insertion(input)
		assert.Equalf(t, expected, sorted, "Insertion sort failed")
	}
}
