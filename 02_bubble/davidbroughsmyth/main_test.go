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
	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
	assert.Equal(t, actual, expected, "Unexpected output from main()")
}

var tests = map[string]struct {
	input []int
	want  []int
}{
	"Reverse Sort":            {input: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	"Already Sorted":          {input: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
	"All same":                {input: []int{8, 8, 8, 8, 8}, want: []int{8, 8, 8, 8, 8}},
	"All zeros":               {input: []int{0, 0, 0, 0, 0}, want: []int{0, 0, 0, 0, 0}},
	"Empty":                   {input: []int{}, want: []int{}},
	"Negative":                {input: []int{-1, -3, -5, -2, -4}, want: []int{-5, -4, -3, -2, -1}},
	"Mixed Neg & Pos numbers": {input: []int{4, -1, -3, 0, -2, 3}, want: []int{-3, -2, -1, 0, 3, 4}},
}

func TestBubble(t *testing.T) {
	for name, test := range tests {
		testList := test
		t.Run(name, func(t *testing.T) {
			got := bubble(testList.input)
			assert.Equal(t, testList.want, got)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for name, test := range tests {
		testList := test
		t.Run(name, func(t *testing.T) {
			got := insertionSort(testList.input)
			assert.Equal(t, testList.want, got)
		})
	}
}
