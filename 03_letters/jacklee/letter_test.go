package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData = []struct {
	in     string
	letter map[rune]int
	out    []string
}{
	{"", map[rune]int{}, []string{}},
	{"aba", map[rune]int{'a': 2, 'b': 1}, []string{"a:2", "b:1"}},
	{"aaAAbb QQa ", map[rune]int{'a': 3, 'A': 2, 'b': 2, 'Q': 2, ' ': 2}, []string{" :2", "A:2", "Q:2", "a:3", "b:2"}},
	{"ああ你你  ", map[rune]int{'あ': 2, '你': 2, ' ': 2}, []string{" :2", "あ:2", "你:2"}},
}

func TestMainOutput(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//When
	main()

	//Then
	expected := "a:2\nb:1\n"
	actual := buf.String()
	r.Equal(expected, actual, "Unexpected output")
}

func TestLetters(t *testing.T) {
	for _, d := range testData {
		// When
		l := letters(d.in)

		// Then
		require.Equal(t, d.letter, l, "unexpected output")
	}
}

func TestSortLetters(t *testing.T) {
	for _, d := range testData {
		// When
		s := sortLetters(d.letter)

		// Then
		require.Equal(t, d.out, s, "unexpected output")
	}
}
