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
	expected := `a : 2
b : 1
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestCountLettersOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := map[rune]int{104: 1, 97: 2, 110: 1, 105: 1, 115: 1}
	actual := countLetters("anisha")
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestCountLettersOutputWithSpace(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := map[rune]int{' ': 2, 'n': 1, 'i': 1, 's': 1, 'h': 1}
	actual := countLetters(" nish ")
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestCountLettersOutputWithSpecialChar(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := map[rune]int{'@': 1, 'n': 1, 'i': 1, '$': 1, 'h': 1, 'a': 1}
	actual := countLetters("@ni$ha")
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestSortLettersOutputWithSpecialChar(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := []string{"$ : 1", "@ : 1", "a : 1", "h : 1", "i : 1", "n : 1"}
	actual := sortLetters(map[rune]int{'@': 1, 'n': 1, 'i': 1, '$': 1, 'h': 1, 'a': 1})
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestSortLettersOutputWithSpace(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := []string{"  : 2", "h : 1", "i : 1", "n : 1", "s : 1"}
	actual := sortLetters(map[rune]int{' ': 2, 'n': 1, 'i': 1, 's': 1, 'h': 1})
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestSortLettersOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// Then
	expected := []string{"  : 2", "h : 1", "i : 1", "n : 1", "s : 1"}
	actual := sortLetters(map[rune]int{' ': 2, 'n': 1, 'i': 1, 's': 1, 'h': 1})
	r.Equalf(expected, actual, "Unexpected output in main()")
}
