package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected []int
	}{
		"test1":       {input: []int{5, 2, 3, 1}, expected: []int{1, 2, 3, 5}},
		"test2":       {input: []int{-9, 2, -12, 12}, expected: []int{-12, -9, 2, 12}},
		"empty":       {input: []int{}, expected: []int{}},
		"sameval":     {input: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
		"combination": {input: []int{3, 4, 2, -1, 3, 5, 3}, expected: []int{-1, 2, 3, 3, 3, 4, 5}},
		"oneval":      {input: []int{11}, expected: []int{11}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := bubble(test.input)
			assert.Equalf(t, result, test.expected, "got %v want %v", result, test.expected)
		})
	}
}

func TestInsertion(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected []int
	}{
		"empty":       {input: []int{}, expected: []int{}},
		"sameval":     {input: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
		"combination": {input: []int{3, 4, 2, -1, 3, 5, 3}, expected: []int{-1, 2, 3, 3, 3, 4, 5}},
		"oneval":      {input: []int{11}, expected: []int{11}},
		"test1":       {input: []int{5, 2, 3, 1}, expected: []int{1, 2, 3, 5}},
		"test2":       {input: []int{-9, 2, -12, 12}, expected: []int{-12, -9, 2, 12}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := insertion(test.input)
			assert.Equalf(t, result, test.expected, "got %v want %v", result, test.expected)
		})
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	expected := "[1 2 3 5]\n"
	main()
	actual := buf.String()
	assert.Equalf(t, actual, expected, "got %v want %v", actual, expected)
}
