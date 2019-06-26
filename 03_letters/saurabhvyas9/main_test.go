package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "Letter frequency failed.")
}

var testCases = map[string]struct {
	input         string
	outputLetters map[rune]int
	outputSort    []string
}{
	"main": {
		input:         "aba",
		outputLetters: map[rune]int{'a': 2, 'b': 1},
		outputSort:    []string{"a:2", "b:1"},
	},
	"no duplicates": {
		input:         "abcd",
		outputLetters: map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1},
		outputSort:    []string{"a:1", "b:1", "c:1", "d:1"},
	},
	"duplicates with space": {
		input:         " aaaxbyb bb c z cddsd sd sdsds",
		outputLetters: map[rune]int{' ': 7, 'a': 3, 'b': 4, 'c': 2, 'd': 6, 's': 5, 'x': 1, 'y': 1, 'z': 1},
		outputSort:    []string{" :7", "a:3", "b:4", "c:2", "d:6", "s:5", "x:1", "y:1", "z:1"},
	},
	"Alphanumeric": {
		input:         "aa234444",
		outputLetters: map[rune]int{'2': 1, '3': 1, '4': 4, 'a': 2},
		outputSort:    []string{"2:1", "3:1", "4:4", "a:2"},
	},
	"empty": {input: "", outputLetters: map[rune]int{}, outputSort: []string{}},
	"capitals": {
		input:         "ABCD",
		outputLetters: map[rune]int{'A': 1, 'B': 1, 'C': 1, 'D': 1},
		outputSort:    []string{"A:1", "B:1", "C:1", "D:1"},
	},
	"special characters": {
		input: "!@#$%^&*()_:><?({})}{];/.[[,`~_++_)(*&^%$#@!~]",
		outputLetters: map[rune]int{'!': 2, '#': 2, '$': 2, '%': 2, '&': 2, '(': 3, ')': 3, '*': 2, '+': 2,
			',': 1, '.': 1, '/': 1, ':': 1, ';': 1, '<': 1, '>': 1, '?': 1, '@': 2, '[': 2, ']': 2,
			'^': 2, '_': 3, '`': 1, '{': 2, '}': 2, '~': 2},
		outputSort: []string{"!:2", "#:2", "$:2", "%:2", "&:2", "(:3", "):3", "*:2", "+:2",
			",:1", ".:1", "/:1", "::1", ";:1", "<:1", ">:1", "?:1", "@:2", "[:2", "]:2",
			"^:2", "_:3", "`:1", "{:2", "}:2", "~:2"},
	},
	"foreign languages": {
		input:         "你好áÁこんにちは",
		outputLetters: map[rune]int{'Á': 1, 'á': 1, 'こ': 1, 'ち': 1, 'に': 1, 'は': 1, 'ん': 1, '你': 1, '好': 1},
		outputSort:    []string{"Á:1", "á:1", "こ:1", "ち:1", "に:1", "は:1", "ん:1", "你:1", "好:1"},
	},
	"emoticons": {
		input:         "😃🐻🙏🔥😂😂🤔😂🔥😂😂🔥",
		outputLetters: map[rune]int{'🐻': 1, '🔥': 3, '😂': 5, '😃': 1, '🙏': 1, '🤔': 1},
		outputSort:    []string{"🐻:1", "🔥:3", "😂:5", "😃:1", "🙏:1", "🤔:1"},
	},
}

func TestSortLetters(t *testing.T) {
	for caseName, test := range testCases {
		input := test.outputLetters
		expected := test.outputSort
		actual := sortLetters(input)
		assert.Equalf(t, expected, actual, "Sort function failed for TC: %d", caseName)
	}
}

func TestLetters(t *testing.T) {
	for caseName, test := range testCases {
		input := test.input
		expected := test.outputLetters
		actual := letters(input)
		assert.Equalf(t, expected, actual, "Letters function failed for TC: %d", caseName)
	}
}
