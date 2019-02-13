package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testMatrix = []struct {
	input      string
	lettersMap map[rune]int
	sortedKeys []string
}{
	{"", map[rune]int{}, []string{}},

	{"abba",
		map[rune]int{('a'): 2, ('b'): 2},
		[]string{"a:2", "b:2"}},

	{"abb a",
		map[rune]int{(' '): 1, ('a'): 2, ('b'): 2},
		[]string{" :1", "a:2", "b:2"}},

	{"£€€€§‡®😎",
		map[rune]int{('£'): 1, ('‡'): 1, ('®'): 1, ('€'): 3, ('§'): 1, ('😎'): 1},
		[]string{"£:1", "§:1", "®:1", "‡:1", "€:3", "😎:1"}}}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `a:2
b:1`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestLetters(t *testing.T) {
	// Given
	r := require.New(t)

	for _, testData := range testMatrix {
		letters := letters(testData.input)
		r.Equal(testData.lettersMap, letters)
	}
}

func TestSortLetters(t *testing.T) {
	// Given
	r := require.New(t)

	for _, testData := range testMatrix {
		sortedKeys := sortLetters(testData.lettersMap)
		r.EqualValues(testData.sortedKeys, sortedKeys)
	}
}
