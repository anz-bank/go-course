package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := `a:2
b:1
`
	actual := buf.String()
	assert.Equalf(t, expected, actual, "got %v want %v", actual, expected)
}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"normal":  {input: "blahbla", expected: map[rune]int{'a': 2, 'b': 2, 'h': 1, 'l': 2}},
		"empty":   {input: "", expected: map[rune]int{}},
		"special": {input: "!@##$", expected: map[rune]int{'!': 1, '@': 1, '#': 2, '$': 1}},
		"emoji":   {input: "😀😎", expected: map[rune]int{'😀': 1, '😎': 1}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := letters(test.input)
			assert.Equalf(t, test.expected, result, "got %v want %v", result, test.expected)
		})
	}
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		expected []string
		input    map[rune]int
	}{
		"normal":  {expected: []string{"k:2", "l:2", "o:1", "z:2"}, input: map[rune]int{'k': 2, 'l': 2, 'z': 2, 'o': 1}},
		"empty":   {expected: []string{}, input: map[rune]int{}},
		"special": {expected: []string{"*:2", "@:3", "^:4"}, input: map[rune]int{'^': 4, '*': 2, '@': 3}},
		"emoji":   {expected: []string{"😀:10", "😎:13"}, input: map[rune]int{'😎': 13, '😀': 10}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := sortLetters(test.input)
			assert.Equalf(t, test.expected, result, "got %v want %v", result, test.expected)
		})
	}
}
