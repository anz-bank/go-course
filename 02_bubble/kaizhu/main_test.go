package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestBubble(t *testing.T) {
	type test struct {
		input    []int
		expected []int
	}

	tests := []test{
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{2, 1}, expected: []int{1, 2}},
		{input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for _, test := range tests {
		actual := bubble(test.input)

		if (actual == nil) != (test.expected == nil) {
			t.Errorf("Unexpected output in bubble()\ninput: %d\nexpected: %q\nactual: %q",
				test.input, test.expected, actual)
		}

		if len(actual) != len(test.expected) {
			t.Errorf("Unexpected output in bubble()\ninput: %d\nexpected: %q\nactual: %q",
				test.input, test.expected, actual)
		}

		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("Unexpected output in bubble()\ninput: %d\nexpected: %q\nactual: %q",
					test.input, test.expected, actual)
			}
		}

	}
}
