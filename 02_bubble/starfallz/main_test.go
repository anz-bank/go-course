package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of bubble([]int{3, 2, 1, 5}) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote("[1 2 3 5]")
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestBubbleFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expected    []int
	}{
		{"Test bubble sorting from already sorted array", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"Test bubble sorting with negative numbers", []int{-1, -2, -3, -4}, []int{-4, -3, -2, -1}},
		{"Test bubble sorting with negative numbers", []int{1, -1, -2, 3, 2, -3, -4, 0, 4},
			[]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}},
		{"Test bubble sorting sorting with Zero", []int{0}, []int{0}},
		{"Test bubble sorting with nil", nil, nil},
	}

	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected

		t.Run(testCase.description, func(t *testing.T) {
			result := bubble(input)

			for i, actual := range result {
				if expected[i] != actual {
					t.Errorf("Unexpected output, expected: %d, actual: %d", expected[i], actual)
				}
			}
		})
	}
}
