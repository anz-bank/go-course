package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	r := require.New(t)
	// Given
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
	letterMap := letters("lokili")
	expected := map[rune]int{'i': 2, 'k': 1, 'l': 2, 'o': 1}
	r.Equalf(expected, letterMap, "Unexpected output in main()")
}

func TestSortLetters(t *testing.T) {
	r := require.New(t)
	input := map[rune]int{'d': 3, 'b': 10, 'z': 7}
	sortedLetters := sortLetters(input)
	expected := []string{"b:10", "d:3", "z:7"}
	r.Equalf(expected, sortedLetters, "Unexpected output in main()")
}

func TestLettersWithSpecialCharacters(t *testing.T) {
	r := require.New(t)
	letterMap := letters("我\\/H&")
	expected := map[rune]int{'我': 1, '\\': 1, '/': 1, '&': 1, 'H': 1}
	r.Equalf(expected, letterMap, "Unexpected output in main()")
}
