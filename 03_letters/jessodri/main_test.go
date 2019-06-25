package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[ :5 !:1 T:1 a:2 f:1 g:2 h:1 i:4 m:1 n:2 o:1 p:2 r:2 s:3 t:1]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestSortLetters(t *testing.T) {
	testCases := []struct {
		description string
		letterFreq  map[rune]int
		expected    []string
	}{
		{
			"happy path",
			map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5},
			[]string{"a:1", "b:2", "c:3", "d:4", "e:5"},
		},
		{
			"characters and empty spaces",
			map[rune]int{' ': 4, '!': 5, '&': 1, 'b': 2, 'c': 3},
			[]string{" :4", "!:5", "&:1", "b:2", "c:3"},
		},
		{
			"empty",
			map[rune]int{},
			[]string{},
		},
	}

	for _, tc := range testCases {
		expected := tc.expected
		actual := sortLetters(tc.letterFreq)
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, actual, expected, "got %v but wanted %v", actual, expected)
		})
	}
}

func TestLetters(t *testing.T) {
	testCases := []struct {
		description string
		letters     string
		expected    map[rune]int
	}{
		{
			"happy path",
			"accdceee",
			map[rune]int{'a': 1, 'c': 3, 'd': 1, 'e': 3},
		},
		{
			"characters and empty spaces",
			"aa ! && cc",
			map[rune]int{' ': 3, 'a': 2, 'c': 2, '!': 1, '&': 2},
		},
		{
			"empty",
			"",
			map[rune]int{},
		},
	}

	for _, tc := range testCases {
		expected := tc.expected
		actual := letters(tc.letters)
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, actual, expected, "got %v but wanted %v", actual, expected)
		})
	}
}
