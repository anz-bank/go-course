package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		arg      string
		expected map[rune]int
	}{
		"Basic": {
			arg:      "Hello",
			expected: map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1},
		},
		"Emoji": {
			arg:      "😀💩💩💩🤡🤡",
			expected: map[rune]int{'😀': 1, '💩': 3, '🤡': 2},
		},
		"Case Sensitive": {
			arg:      "Abba",
			expected: map[rune]int{'A': 1, 'b': 2, 'a': 1},
		},
		"Empty": {
			arg:      "",
			expected: map[rune]int{},
		},
	}

	for testName, tC := range testCases {
		testCase := tC
		t.Run(testName, func(t *testing.T) {
			actual := letters(testCase.arg)
			expected := testCase.expected
			assert.Equal(t, expected, actual)
		})

	}
}
func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		arg      map[rune]int
		expected []string
	}{
		"Basic": {
			arg:      map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1},
			expected: []string{"l:2", "H:1", "e:1", "o:1"},
		},
		"Emoji": {
			arg:      map[rune]int{'😀': 1, '💩': 3, '🤡': 2},
			expected: []string{"💩:3", "🤡:2", "😀:1"},
		},
		"Case Sensitive": {
			arg:      map[rune]int{'A': 1, 'b': 2, 'a': 1},
			expected: []string{"b:2", "A:1", "a:1"},
		},
		"Empty": {
			arg:      map[rune]int{},
			expected: []string{},
		},
	}

	for testName, tC := range testCases {
		testCase := tC
		t.Run(testName, func(t *testing.T) {
			actual := sortLetters(testCase.arg)
			expected := testCase.expected
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	actual := buf.String()
	expected := "a:2\nb:1\n"
	assert.Equal(t, expected, actual)
}
