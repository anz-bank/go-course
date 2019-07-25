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
	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"No params": {input: "", expected: map[rune]int{}},
		"1 char":    {input: "a", expected: map[rune]int{97: 1}},
		"2 chars":   {input: "a0", expected: map[rune]int{48: 1, 97: 1}},
		"2 same":    {input: "aa", expected: map[rune]int{97: 2}},
		"utf8 char": {input: "£a£0", expected: map[rune]int{48: 1, 97: 1, 163: 2}},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := letters(test.input)
			expected := test.expected
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"No params": {input: map[rune]int{}, expected: []string{}},
		"1 char":    {input: map[rune]int{97: 1}, expected: []string{"a:1"}},
		"2 chars":   {input: map[rune]int{97: 1, 48: 1}, expected: []string{"0:1", "a:1"}},
		"2 same":    {input: map[rune]int{97: 2}, expected: []string{"a:2"}},
		"utf8 char": {input: map[rune]int{97: 1, 163: 2}, expected: []string{"a:1", "£:2"}},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := sortLetters(test.input)
			expected := test.expected
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}
