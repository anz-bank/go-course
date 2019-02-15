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
	expected := `a:2
b:1`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestLetters(t *testing.T) {
	r := require.New(t)
	letterMap := letters("vinila")
	expected := map[rune]int{'a': 1, 'i': 2, 'l': 1, 'n': 1, 'v': 1}
	r.Equalf(expected, letterMap, "Unexpected output in main()")
}

func TestSortLetters(t *testing.T) {
	r := require.New(t)
	input := map[rune]int{'v': 4, 'i': 3, 'n': 1, 'l': 1}
	sortedKeys := sortLetters(input)
	expected := []string{"i:3", "l:1", "n:1", "v:4"}
	r.Equalf(expected, sortedKeys, "Unexpected output in main()")
}

func TestLettersWithSpecialCharecters(t *testing.T) {
	r := require.New(t)
	letterMap := letters("we 🖤 go")
	expected := map[rune]int{'w': 1, 'e': 1, ' ': 2, '🖤': 1, 'g': 1, 'o': 1}
	r.Equalf(expected, letterMap, "Unexpected output in main()")
}
