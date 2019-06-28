package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Edge Case - Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Edge Case - Singular slice",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Edge Case - Smallest slice that requires sorting",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Best Case",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Average Case",
			input:    []int{12, 2, 1, 6, 9, 3, 10},
			expected: []int{1, 2, 3, 6, 9, 10, 12},
		},
		{
			name:     "Worst Case",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "With Duplicates",
			input:    []int{3, 2, 3, 6, 8, 6, 9},
			expected: []int{2, 3, 3, 6, 6, 8, 9},
		},
	}

	sortFns := map[string]func([]int) []int{
		"Bubble Sort":    bubble,
		"Insertion Sort": insertion,
	}

	for fname, fn := range sortFns {
		fn := fn
		for _, testCase := range cases {
			testCase := testCase
			t.Run(fname+":"+testCase.name, func(t *testing.T) {
				assert.Equal(t, testCase.expected, fn(testCase.input))
			})
		}
	}
}

func TestBubbleDoesNotTouchInput(t *testing.T) {
	ns := []int{3, 2, 1}
	bubble(ns)
	assert.Equal(t, []int{3, 2, 1}, ns)
}

func TestInsertionDoesNotTouchInput(t *testing.T) {
	ns := []int{3, 2, 1}
	insertion(ns)
	assert.Equal(t, []int{3, 2, 1}, ns)
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[1 2 3]\n[1 2 3]\n", buf.String())
}
