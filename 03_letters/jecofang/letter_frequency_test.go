package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "a:2\nb:1\n"
	actual := buf.String()
	r.Equal(expected, actual, "Unexpected output in main()")
}

var testData = []struct {
	in      string
	letters map[rune]int
	sorted  []string
}{
	{"", map[rune]int{}, []string{}},
	{"aba", map[rune]int{'a': 2, 'b': 1}, []string{"a:2", "b:1"}},
	{"abaAbA", map[rune]int{'a': 2, 'b': 2, 'A': 2}, []string{"A:2", "a:2", "b:2"}},
	{"ab aAbA", map[rune]int{'a': 2, 'b': 2, ' ': 1, 'A': 2}, []string{" :1", "A:2", "a:2", "b:2"}},
	{"ab aAbA21",
		map[rune]int{'a': 2, 'b': 2, ' ': 1, 'A': 2, '2': 1, '1': 1},
		[]string{" :1", "1:1", "2:1", "A:2", "a:2", "b:2"}},
	{"ab aAbA21你 ",
		map[rune]int{'a': 2, 'b': 2, ' ': 2, 'A': 2, '2': 1, '1': 1, '你': 1},
		[]string{" :2", "1:1", "2:1", "A:2", "a:2", "b:2", "你:1"}},
	{"ab,aAbA21你，",
		map[rune]int{'a': 2, 'b': 2, ',': 1, 'A': 2, '2': 1, '1': 1, '你': 1, '，': 1},
		[]string{",:1", "1:1", "2:1", "A:2", "a:2", "b:2", "你:1", "，:1"}},
}

func TestLetters(t *testing.T) {
	for _, d := range testData {
		// When
		l := letters(d.in)

		// Then
		require.Equal(t, d.letters, l, "unexpected output")
	}
}

func TestSortLetters(t *testing.T) {
	for _, d := range testData {
		// When
		s := sortLetters(d.letters)

		// Then
		require.Equal(t, d.sorted, s, "unexpected output")
	}
}
