package main

import (
	"bytes"

	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}

}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"Happy case": {"aba", map[rune]int{97: 2, 98: 1}},
		"Map Contains numbers": {"abadddcc311a", map[rune]int{'1': 2, '3': 1, 97: 3, 98: 1,
			99: 2, 100: 3}},
		"Empty map should return empty slice": {"", map[rune]int{}},
		"Map Contains emoji or even non-english characters ": {"😀😀😀🤓🤓哈哈哈你大爷的", map[rune]int{
			'你': 1, '哈': 3, '大': 1, '爷': 1, '的': 1, '😀': 3, '🤓': 2}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			expected := test.expected
			actual := lettersFreq(test.input)
			if !assert.Equal(t, test.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"Happy case": {map[rune]int{97: 2, 98: 1}, []string{"a:2", "b:1"}},
		"Map Contains numbers": {map[rune]int{1: 2, 3: 1, 97: 3, 98: 1, 99: 2, 100: 3},
			[]string{"\x01:2", "\x03:1", "a:3", "b:1", "c:2", "d:3"}},
		"Empty map should return empty slice": {map[rune]int{}, []string{}},
		"Map Contains uppercase letters": {map[rune]int{97: 3, 98: 1, 99: 8, 100: 3, 67: 1},
			[]string{"C:1", "a:3", "b:1", "c:8", "d:3"}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			expected := test.expected
			actual := sortLetters(test.input)
			if !assert.Equal(t, test.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
			}
		})
	}
}
