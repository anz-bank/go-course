package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"polish": {input: "dzień dobry",
			expected: map[rune]int{'y': 1, ' ': 1, 'd': 2, 'z': 1, 'i': 1, 'e': 1, 'ń': 1, 'o': 1, 'b': 1, 'r': 1}},
		"russian": {input: "доброе утро",
			expected: map[rune]int{'о': 3, 'р': 2, 'т': 1, 'у': 1, ' ': 1, 'е': 1, 'б': 1, 'д': 1}},
		"english": {input: "good morning",
			expected: map[rune]int{'o': 3, 'g': 2, 'd': 1, 'm': 1, 'r': 1, 'n': 2, 'i': 1, ' ': 1}},
		"chineese": {input: "早上好",
			expected: map[rune]int{'早': 1, '上': 1, '好': 1}},
	}

	for name, test := range testCases {
		t.Run(name, helperTestLetter(test.input, test.expected))
	}
}

func helperTestLetter(input string, expected map[rune]int) func(*testing.T) {
	return func(t *testing.T) {
		result := letters(input)
		for k, count := range result {
			if _, ok := expected[k]; !ok {
				t.Errorf("Found unexpected key: %c", k)
			}
			assert.Equalf(t, expected[k], count, "Incorrect key")
		}
		assert.Equalf(t, len(expected), len(result), "Length mismatch")

	}
}

func TestLetterFreqStringNegative(t *testing.T) {
	testCases := map[string]struct {
		input       letterFreq
		notExpected string
	}{
		// cyrillic 'а' is different char than latin 'a'
		"mixed": {input: letterFreq{letter: 'а', count: 2}, notExpected: "a:2"},
		"count": {input: letterFreq{letter: 'a', count: 2}, notExpected: "a:3"},
		"key":   {input: letterFreq{letter: 'b', count: 2}, notExpected: "a:2"},
	}
	for name, test := range testCases {
		input := test.input
		notExpected := test.notExpected
		t.Run(name, func(t *testing.T) {
			result := input.String()
			assert.NotEqualf(t, notExpected, result, "Unexpected result")
		})
	}

}

func TestLetterFreqStringPositive(t *testing.T) {
	testCases := map[string]struct {
		input    letterFreq
		expected string
	}{
		"latinum":  {input: letterFreq{letter: 'a', count: 2}, expected: "a:2"},
		"cyrillic": {input: letterFreq{letter: 'а', count: 2}, expected: "а:2"},
	}
	for name, test := range testCases {
		input := test.input
		expected := test.expected
		t.Run(name, func(t *testing.T) {
			result := input.String()
			assert.Equalf(t, expected, result, "Unexpected result")
		})
	}

}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"latinum":  {input: map[rune]int{'c': 2, 'a': 1, 'b': 3}, expected: []string{"b:3", "c:2", "a:1"}},
		"cyrillic": {input: map[rune]int{'в': 2, 'а': 1, 'б': 3}, expected: []string{"б:3", "в:2", "а:1"}},
	}
	for name, test := range testCases {
		t.Run(name, helperTestSortLetters(name, test.input, test.expected))
	}
}

func helperTestSortLetters(name string, input map[rune]int, expected []string) func(*testing.T) {
	return func(t *testing.T) {
		result := sortLetters(input)
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Test: %s - Expect: %s got: %s", name, expected[i], v)
			}
		}
		assert.Equalf(t, len(expected), len(result), "Length mismatch")
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}
