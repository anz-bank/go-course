package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[1 2 3 5]\n[1 2 3 5]"
	got := buf.String()

	t.Run("Main function", func(t *testing.T) {
		if expected != got {
			t.Errorf("\nExpected: %#v\nGot:      %#v", expected, got)
		}
	})
}

var cases = map[string]struct {
	input    []int
	expected []int
}{
	"Test one": {
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	"Test two": {
		input:    []int{1, 1, 2, 0, 1, 2},
		expected: []int{0, 1, 1, 1, 2, 2},
	},
	"Test three": {
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	"Negative": {
		input:    []int{-1, -2, -3, -4, -5},
		expected: []int{-5, -4, -3, -2, -1},
	},
	"Negative and positive": {
		input:    []int{-1, -2, -3, 2, 1},
		expected: []int{-3, -2, -1, 1, 2},
	},
	"Empty": {
		input:    []int{},
		expected: []int{},
	},
	"Single": {
		input:    []int{5},
		expected: []int{5},
	},
}

func TestBubble(t *testing.T) {
	for name, c := range cases {
		got, expected := bubble(c.input), c.expected
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("\nExpected: %d\nGot:      %d", expected, got)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	for name, c := range cases {
		got, expected := insertion(c.input), c.expected
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("\nExpected: %d\nGot:      %d", expected, got)
			}
		})
	}
}
