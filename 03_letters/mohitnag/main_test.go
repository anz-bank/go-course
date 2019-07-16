package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	type testCases struct {
		input    string
		expected map[rune]int
	}
	testData := map[string]testCases{
		"Happy Case": {"eaa!baabq23123!",
			map[rune]int{'!': 2, '1': 1, '2': 2, '3': 2, 'a': 4, 'b': 2, 'e': 1, 'q': 1}},
		"Empty string should return empty map": {"", map[rune]int{}},
		"Unicode string test":                  {"😀😀😀😀🤓🤓", map[rune]int{'😀': 4, '🤓': 2}},
	}

	for scenario, td := range testData {
		td := td
		t.Run(scenario, func(t *testing.T) {
			letterMap := letters(td.input)
			assert.Equal(t, td.expected, letterMap)
		})
	}
}
func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "a:2\nb:1\n"
	actual := buf.String()
	assert.Equal(expected, actual)
}

func TestSortLetters(t *testing.T) {
	type testCases struct {
		input    map[rune]int
		expected []string
	}
	testData := map[string]testCases{

		"Happy Case": {map[rune]int{'!': 2, '1': 1, '2': 2, '3': 2, 'a': 4, 'b': 2, 'e': 1, 'q': 1},
			[]string{"!:2", "1:1", "2:2", "3:2", "a:4", "b:2", "e:1", "q:1"}},
		"Empty map should return empty slice": {map[rune]int{}, []string{}},
	}
	for scenario, td := range testData {
		td := td
		t.Run(scenario, func(t *testing.T) {
			sortedLetter := sortLetters(td.input)
			assert.Equal(t, td.expected, sortedLetter)
		})
	}
}
