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

	expected := strconv.Quote("[1 2 3 4 5 43 67]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("actual: %s does not match expected: %s", actual, expected)
	}
}

func TestSorting(t *testing.T) {
	var testCases = []struct {
		description string
		unsorted    []int
		sorted      []int
	}{
		{"happy path", []int{42, 2, 34, 8, 5, 19, 11}, []int{2, 5, 8, 11, 19, 34, 42}},
		{"duplicate numbers", []int{88, 88, 2, 56, 14, 88, 23, 2}, []int{2, 2, 14, 23, 56, 88, 88, 88}},
		{"negative numbers", []int{1, -3, 67, -48, 4, 6, 0}, []int{-48, -3, 0, 1, 4, 6, 67}},
		{"empty array", []int{}, []int{}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := bubble(tc.unsorted)
			expected := tc.sorted
			if !assert.ElementsMatch(t, expected, actual) {
				t.Errorf("array was not sorted,\nexpected: %d\nactual: %d", expected, actual)
			}
		})
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := insertion(tc.unsorted)
			expected := tc.sorted
			if !assert.ElementsMatch(t, expected, actual) {
				t.Errorf("array was not sorted,\nexpected: %d\nactual: %d", expected, actual)
			}
		})
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := quicksort(tc.unsorted)
			expected := tc.sorted
			if !assert.ElementsMatch(t, expected, actual) {
				t.Errorf("array was not sorted,\nexpected: %d\nactual: %d", expected, actual)
			}
		})
	}
}
