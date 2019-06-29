package main

import (
	"bytes"
	"reflect"
	"testing"
)

var testCases = map[string]struct {
	input    []int
	expected []int
}{
	"Basic":    {input: []int{3, 2, 1, 5}, expected: []int{1, 2, 3, 5}},
	"Empty":    {input: []int{}, expected: []int{}},
	"Equal":    {input: []int{7, 7}, expected: []int{7, 7}},
	"Large":    {input: []int{34621, 78234, 95324, 10129}, expected: []int{10129, 34621, 78234, 95324}},
	"Negative": {input: []int{2, -3, 9, -8, -1, 0}, expected: []int{-8, -3, -1, 0, 2, 9}},
	"Reverse":  {input: []int{89, 63, 58, 42, 24, 13, 9}, expected: []int{9, 13, 24, 42, 58, 63, 89}},
}

func TestBubble(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		expected := test.expected
		t.Run(name, func(t *testing.T) {
			actual := bubble(input)
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("Unexpected output in bubble():\n\texpected: %v\n\tgot: %v", expected, actual)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		expected := test.expected
		t.Run(name, func(t *testing.T) {
			actual := insertion(input)
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("Unexpected output in insertion():\n\texpected: %v\n\tgot: %v", expected, actual)
			}
		})
	}
}

func TestQuick(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		expected := test.expected
		t.Run(name, func(t *testing.T) {
			actual := quick(input)
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("Unexpected output in quick():\n\texpected: %v\n\tgot: %v", expected, actual)
			}
		})
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "[1 2 3 5]\n"
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in main():\n\texpected %v\n\tgot: %v", expected, actual)
	}
}
