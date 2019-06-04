package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestLetters(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"Empty string": {
			input:    "",
			expected: map[rune]int{}},
		"ASCII string": {
			input: "  ``~8I,dzye[uY6<mCh<n 9Otefp0fX0-@2<C)z)}.go-Hq{n]LX 8uKnRxj (92@08b9P ",
			expected: map[rune]int{
				'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
				' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
				'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
				']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
			}},
		"Rune string": {
			input: "菎哫😄₲৲₯₷$€≷⊇∲ℬ℠Ω🀶🃊🀆🁒🂁ⱭḸṄ☞✤❡쉱🤒🐧췡㼪ν㰗👀뾬✌",
			expected: map[rune]int{
				'㰗': 1, '㼪': 1, '哫': 1, '菎': 1, '뾬': 1, '쉱': 1, '췡': 1, '🀆': 1, '🀶': 1, '🁒': 1,
				'Ω': 1, 'ℬ': 1, '∲': 1, '≷': 1, '⊇': 1, '☞': 1, '✌': 1, '✤': 1, '❡': 1, 'Ɑ': 1,
				'$': 1, 'ν': 1, '৲': 1, 'Ḹ': 1, 'Ṅ': 1, '€': 1, '₯': 1, '₲': 1, '₷': 1, '℠': 1,
				'🂁': 1, '🃊': 1, '🐧': 1, '👀': 1, '😄': 1, '🤒': 1,
			}},
		"String with esc chars": {
			input: "\"Kia ora\"",
			expected: map[rune]int{
				'i': 1, 'o': 1, 'r': 1, ' ': 1, '"': 2, 'K': 1, 'a': 2,
			}},
	}

	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := letters(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				fmt.Println(actual)
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %v", testData.expected, actual)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	tests := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"Empty map": {
			input:    map[rune]int{},
			expected: []string{}},
		"ASCII map": {
			input: map[rune]int{
				'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
				' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
				'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
				']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
			},
			expected: []string{
				" :6", "(:1", "):2", ",:1", "-:2", ".:1", "0:3", "2:2", "6:1", "8:3", "9:3", "<:3", "@:2", "C:2",
				"H:1", "I:1", "K:1", "L:1", "O:1", "P:1", "R:1", "X:2", "Y:1", "[:1", "]:1", "`:2", "b:1", "d:1",
				"e:2", "f:2", "g:1", "h:1", "j:1", "m:1", "n:3", "o:1", "p:1", "q:1", "t:1", "u:2", "x:1", "y:1",
				"z:2", "{:1", "}:1", "~:1",
			}},
		"Rune map": {
			input: map[rune]int{
				'㰗': 1, '㼪': 1, '哫': 1, '菎': 1, '뾬': 1, '쉱': 1, '췡': 1, '🀆': 1, '🀶': 1, '🁒': 1,
				'Ω': 1, 'ℬ': 1, '∲': 1, '≷': 1, '⊇': 1, '☞': 1, '✌': 1, '✤': 1, '❡': 1, 'Ɑ': 1,
				'$': 1, '🃊': 1, '🐧': 1, '👀': 1, '😄': 1, '🤒': 1, '₯': 1, '₲': 1, '₷': 1, '℠': 1,
				'🂁': 1, 'ν': 1, '৲': 1, 'Ḹ': 1, 'Ṅ': 1, '€': 1,
			},
			expected: []string{
				"$:1", "ν:1", "৲:1", "Ḹ:1", "Ṅ:1", "€:1", "₯:1", "₲:1", "₷:1", "℠:1", "Ω:1", "ℬ:1",
				"∲:1", "≷:1", "⊇:1", "☞:1", "✌:1", "✤:1", "❡:1", "Ɑ:1", "㰗:1", "㼪:1", "哫:1", "菎:1", "뾬:1", "쉱:1",
				"췡:1", "🀆:1", "🀶:1", "🁒:1", "🂁:1", "🃊:1", "🐧:1", "👀:1", "😄:1", "🤒:1",
			}},
		"Map with esc chars": {
			input: map[rune]int{
				'i': 1, 'o': 1, 'r': 1, ' ': 1, '"': 2, 'K': 1, 'a': 2,
			},
			expected: []string{
				" :1", "\":2", "K:1", "a:2", "i:1", "o:1", "r:1",
			}},
	}

	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := sortLetters(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %s\nactual: %s", testData.expected, actual)
			}
		})
	}
}
