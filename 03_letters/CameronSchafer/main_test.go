package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := `a:2
b:1
`
	actual := buf.String()

	assert.Equalf(t, expected, actual,
		"Input does not match expected buffer output.")
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input       map[rune]int
		expectedArr []string
	}{
		"2 letters singles [a:1, b:1]": {input: map[rune]int{'a': 1, 'b': 1},
			expectedArr: []string{"a:1", "b:1"},
		},
		"2 letters multiples": {input: map[rune]int{'a': 7, 'b': 7},
			expectedArr: []string{"a:7", "b:7"},
		},
		"1 letter, 1 empty": {input: map[rune]int{' ': 2, 'a': 1},
			expectedArr: []string{" :2", "a:1"},
		},
		"1 letter": {input: map[rune]int{'A': 7},
			expectedArr: []string{"A:7"},
		},
		"3 letters": {input: map[rune]int{'A': 1, 'B': 1, 'c': 2},
			expectedArr: []string{"A:1", "B:1", "c:2"},
		},
		"3 random order": {input: map[rune]int{'B': 1, 'A': 1, 'c': 2},
			expectedArr: []string{"A:1", "B:1", "c:2"},
		},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			resultArr := sortLetters(test.input)
			assert.Equalf(t, test.expectedArr, resultArr,
				"Input does not match expected output.")
		})
	}
}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expectedArr map[rune]int
	}{
		"3 letters": {input: "aba",
			expectedArr: map[rune]int{'a': 2, 'b': 1},
		},
		"letters with spaces": {input: "a a ",
			expectedArr: map[rune]int{' ': 2, 'a': 2},
		},
		"AaAa}{}{}}{}{}{": {input: "AaAa}{}{}}{}{}{",
			expectedArr: map[rune]int{'A': 2, 'a': 2, '{': 5, '}': 6},
		},
		"empty string": {input: "",
			expectedArr: map[rune]int{},
		},
		"only spaces": {input: "    ",
			expectedArr: map[rune]int{' ': 4},
		},
		"special characters": {input: "无😊👍👍",
			expectedArr: map[rune]int{'👍': 2, '无': 1, '😊': 1},
		},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			resultArr := letters(test.input)
			assert.Equalf(t, test.expectedArr, resultArr,
				"Input does not match expected output.")
		})
	}
}
