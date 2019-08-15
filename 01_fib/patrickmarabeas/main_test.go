package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := fmt.Sprintf("%s\n", "0\n1\n1\n2\n3\n5\n8")
	got := buf.String()

	t.Run("Main function", func(t *testing.T) {
		if expected != got {
			t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
		}
	})
}

func TestFibError(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(100)

	expected := "fib: 100 is outside of limits of -92..92"
	got := buf.String()

	t.Run("Main function", func(t *testing.T) {
		if expected != got {
			t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
		}
	})
}

func TestCalculateFibSequence(t *testing.T) {
	var cases = map[string]struct {
		input     int
		expected1 []int
		expected2 error
	}{
		"Zero": {
			input:     0,
			expected1: nil,
			expected2: nil,
		},
		"Input of 1": {
			input:     1,
			expected1: []int{0},
			expected2: nil,
		},
		"Input of negative 1": {
			input:     -1,
			expected1: []int{0},
			expected2: nil,
		},
		"Input of 3": {
			input:     3,
			expected1: []int{0, 1, 1},
			expected2: nil,
		},
		"Input of 7": {
			input:     7,
			expected1: []int{0, 1, 1, 2, 3, 5, 8},
			expected2: nil,
		},
		"Input of negative 3": {
			input:     -3,
			expected1: []int{0, -1, 1},
			expected2: nil,
		},
		"Input of negative 7": {
			input:     -7,
			expected1: []int{0, -1, 1, -2, 3, -5, 8},
			expected2: nil,
		},
		"Out of range - positive": {
			input:     100,
			expected1: nil,
			expected2: errors.New("fib: 100 is outside of limits of -92..92"),
		},
		"Out of range - negative": {
			input:     -100,
			expected1: nil,
			expected2: errors.New("fib: -100 is outside of limits of -92..92"),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			got1, got2 := calculateFibSequence(c.input)
			expected1 := c.expected1
			expected2 := c.expected2
			if !reflect.DeepEqual(got1, expected1) || !reflect.DeepEqual(got2, expected2) {
				t.Errorf("\nExpected: %d, %s\nGot:      %d, %s", expected1, expected2, got1, got2)
			}
		})
	}
}
