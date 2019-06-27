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

func TestLettersFunc(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected map[rune]int
	}

	tests := []test{
		{name: "empty string", input: "", expected: map[rune]int{}}, // expect empty map
		{name: "ASCII string", input: "  ``~8I,dzye[uY6<mCh<n 9Otefp0fX0-@2<C)z)}.go-Hq{n]LX 8uKnRxj (92@08b9P ",
			expected: map[rune]int{
				'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
				' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
				'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
				']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
			}},
		{name: "special characters", input: "无马e😊e马👍👍a👌😒",
			expected: map[rune]int{'a': 1, '马': 2, '👍': 2, '👌': 1, 'e': 2, '无': 1, '😊': 1, '😒': 1}},
		{name: "correct counts", input: "abbcccddddeeeee", expected: map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}},
	}

	for _, testCase := range tests {
		actual := letters(testCase.input)
		fmt.Println(testCase.name, "expected: ", testCase.expected)
		fmt.Println(testCase.name, "actual: ", actual)
		if !reflect.DeepEqual(testCase.expected, actual) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestSortLettersFunction(t *testing.T) {
	type test struct {
		name     string
		input    map[rune]int
		expected []string
	}

	tests := []test{
		{name: "empty map", input: map[rune]int{}, expected: []string{}},
		{name: "ASCII map", input: map[rune]int{
			'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
			' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
			'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
			']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
		}, expected: []string{
			" :6", "(:1", "):2", ",:1", "-:2", ".:1", "0:3", "2:2", "6:1", "8:3", "9:3", "<:3", "@:2", "C:2",
			"H:1", "I:1", "K:1", "L:1", "O:1", "P:1", "R:1", "X:2", "Y:1", "[:1", "]:1", "`:2", "b:1", "d:1",
			"e:2", "f:2", "g:1", "h:1", "j:1", "m:1", "n:3", "o:1", "p:1", "q:1", "t:1", "u:2", "x:1", "y:1",
			"z:2", "{:1", "}:1", "~:1",
		}},
		{name: "special characters: Han Script", input: map[rune]int{'⽖': 1, '⽉': 2, '⽏': 3, '⽐': 2, '⽕': 1},
			expected: []string{
				"⽉:2", "⽏:3", "⽐:2", "⽕:1", "⽖:1",
			}},
		{name: "special characters: Emojis", input: map[rune]int{'🏦': 1, '🔫': 4, '💰': 1, '🚗': 1, '😬': 1,
			'🚓': 7, '😱': 1, '🚒': 1, '🚑': 1, '😨': 1, '😢': 7, '😰': 1,
			'🗯': 1, '🏛': 1, '⏸': 1, '🔒': 1, '👮': 1}, expected: []string{
			"⏸:1", "🏛:1", "🏦:1", "👮:1", "💰:1", "🔒:1", "🔫:4",
			"🗯:1", "😢:7", "😨:1", "😬:1", "😰:1", "😱:1", "🚑:1",
			"🚒:1", "🚓:7", "🚗:1"},
		}}

	for _, testCase := range tests {
		actual := sortLetters(testCase.input)
		fmt.Println(testCase.name, "expected: ", testCase.expected)
		fmt.Println(testCase.name, "actual: ", actual)
		if !reflect.DeepEqual(testCase.expected, actual) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
