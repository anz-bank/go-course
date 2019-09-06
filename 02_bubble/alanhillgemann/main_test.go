package main

import (
	"bytes"

	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("[1 2 3 5]")
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual)
}

var testCases = map[string]struct {
	input    []int
	expected []int
}{
	"No params": {input: []int{}, expected: []int{}},
	"1 int":     {input: []int{1}, expected: []int{1}},
	"Neg int":   {input: []int{2, 3, -1, 0}, expected: []int{-1, 0, 2, 3}},
	"Same ints": {input: []int{2, 1, 1}, expected: []int{1, 1, 2}},
	"Sorted":    {input: []int{0, 1, 2}, expected: []int{0, 1, 2}},
}

func TestBubbleSort(t *testing.T) {
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := bubble(test.input)
			expected := test.expected
			assert.Equal(t, expected, actual)

		})
	}
}
