package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf
	main()
	expected := `a:2
b:1`
	actual := buf.String()
	if expected != actual {
		t.Errorf(expected, actual, "Unexpected output in main()")
	}
}

var lettersCases = map[string]struct {
	input    string
	expected map[rune]int
}{
	"Standard": {input: "aba", expected: map[rune]int{'a': 2, 'b': 1}},
	"Same":     {input: "cccccccccccccccccccc", expected: map[rune]int{'c': 20}},
	"Empty":    {input: "", expected: map[rune]int{}},
	"Return":   {input: "ab\na\n", expected: map[rune]int{'\n': 2, 'a': 2, 'b': 1}},
	"Slash":    {input: "ab/na/n", expected: map[rune]int{'/': 2, 'a': 2, 'b': 1, 'n': 2}},
	"LongMixed": {input: "zabaaarzzzASYbcd@&^!*JHDJKHHGFGD^!*Dar",
		expected: map[rune]int{'!': 2, '&': 1, '*': 2, '@': 1, 'A': 1, 'D': 3, 'F': 1, 'G': 2, 'H': 3, 'J': 2, 'K': 1, 'S': 1,
			'Y': 1, '^': 2, 'a': 5, 'b': 2, 'c': 1, 'd': 1, 'r': 2, 'z': 4}},
	"Punctuation": {input: ".:,!;'?.,:!'",
		expected: map[rune]int{'!': 2, '\'': 2, ',': 2, '.': 2, ':': 2, ';': 1, '?': 1}},
	"Raw": {input: `a:2b:1`, expected: map[rune]int{'1': 1, '2': 1, ':': 2, 'a': 1, 'b': 1}},
	"RawReturn": {input: `a:2
	b:1`, expected: map[rune]int{'	': 1, '\n': 1, '1': 1, '2': 1, ':': 2, 'a': 1, 'b': 1}},
	"Numeric":      {input: "11235", expected: map[rune]int{'1': 2, '2': 1, '3': 1, '5': 1}},
	"EscapedQuote": {input: "a\":2\"b:\"1", expected: map[rune]int{'"': 3, '1': 1, '2': 1, ':': 2, 'a': 1, 'b': 1}},
	"RawQuote":     {input: `a":2"b:"1`, expected: map[rune]int{'"': 3, '1': 1, '2': 1, ':': 2, 'a': 1, 'b': 1}},
	"MultiByteMix": {input: "日本語A⌘дтB語A⌘д",
		expected: map[rune]int{'A': 2, 'B': 1, 'д': 2, 'т': 1, '⌘': 2, '日': 1, '本': 1, '語': 2}},
}

func TestLetters(t *testing.T) {
	for name, tc := range lettersCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := letters(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Letters function failed. Input: %v, Actual: %v, Expected: %v", tc.input, actual, tc.expected)
			}
		})
	}
}

var sortLettersCases = map[string]struct {
	input    map[rune]int
	expected []string
}{
	"Standard": {input: map[rune]int{'b': 1, 'a': 2}, expected: []string{"a:2", "b:1"}},
	"Empty":    {input: map[rune]int{}, expected: []string{}},
	"Long": {input: map[rune]int{'J': 2, 'K': 1, 'S': 1, 'Y': 1, '^': 2, 'a': 5, 'b': 2, 'c': 1, 'd': 1,
		'r': 2, 'z': 4, '!': 2, '&': 1, '*': 2, '@': 1, 'A': 1, 'D': 3, 'F': 1, 'G': 2, 'H': 3},
		expected: []string{"!:2", "&:1", "*:2", "@:1", "A:1", "D:3", "F:1", "G:2", "H:3", "J:2", "K:1",
			"S:1", "Y:1", "^:2", "a:5", "b:2", "c:1", "d:1", "r:2", "z:4"}},
	"Strange": {input: map[rune]int{'о': 3, 'р': 2, 'т': 1, 'у': 1, ' ': 1, 'е': 1, 'б': 1, 'д': 1},
		expected: []string{" :1", "б:1", "д:1", "е:1", "о:3", "р:2", "т:1", "у:1"}},
}

func TestSortLetters(t *testing.T) {
	for name, tc := range sortLettersCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := sortLetters(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Sort Letters function failed. Input: %v, Actual: %v, Expected: %v", tc.input, actual, tc.expected)
			}
		})
	}
}
