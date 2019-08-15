package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	algorithms := map[string]func([]int) []int{
		"Bubble":    bubble,
		"Insertion": insertion,
		"Quick":     quick,
	}

	testCases := map[string]struct {
		arg      []int
		expected []int
	}{
		"Basic": {
			arg:      []int{3, 2, 1, 5, 3, 2, 1, 7, 6},
			expected: []int{1, 1, 2, 2, 3, 3, 5, 6, 7},
		},
		"NegativeValues": {
			arg:      []int{3, -2, 1, 5, -3, 2, 1, -7},
			expected: []int{-7, -3, -2, 1, 1, 2, 3, 5},
		},
		"SingleEntry": {
			arg:      []int{1},
			expected: []int{1},
		},
		"Empty": {
			arg:      []int{},
			expected: []int{},
		},
	}

	for algorithmName, algorithmImpl := range algorithms {
		for scenarioName, tC := range testCases {
			testName := fmt.Sprintf("%v/%v", algorithmName, scenarioName)
			alg := algorithmImpl
			arg := tC.arg
			expected := tC.expected
			t.Run(testName, func(t *testing.T) {
				actual := alg(arg)
				expected := expected
				assert.Equal(t, expected, actual)
			})
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	actual := buf.String()
	expected := "[1 2 3 5]\n[1 2 3 5]\n[1 2 3 5]\n"
	assert.Equal(t, expected, actual)
}

func TestNoMutation(t *testing.T) {
	algorithms := []struct {
		desc string
		fn   func([]int) []int
	}{
		{"Bubble", bubble},
		{"Insertion", insertion},
		{"Quick", quick},
	}
	for _, algorithm := range algorithms {
		fn := algorithm.fn

		t.Run(algorithm.desc, func(t *testing.T) {
			actual := []int{3, 1, 2, 7, -5, 8}
			expected := []int{3, 1, 2, 7, -5, 8}
			fn(actual)
			assert.Equal(t, expected, actual)
		})
	}
}
