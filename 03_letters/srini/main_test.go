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
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	letterMap := letters("abba")

	// Then
	expected := map[rune]int{rune('a'): 2, rune('b'): 2}

	r.Equalf(expected, letterMap, "Unexpected output in main()")
}

func TestLettersWithSpaces(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	letterMap := letters("abb a")

	// Then
	expected := map[rune]int{rune(' '): 1, rune('a'): 2, rune('b'): 2}

	r.Equalf(expected, letterMap, "Unexpected output in main()")
}

func TestLettersWithNonAscii(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	letterMap := letters("£€€€§‡®")

	// Then
	expected := map[rune]int{rune('£'): 1, rune('‡'): 1, rune('®'): 1, rune('€'): 3, rune('§'): 1}

	r.Equalf(expected, letterMap, "Unexpected output in main()")
}

func TestSortLetters(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	input := map[rune]int{rune('x'): 2, rune('b'): 3, rune('t'): 3}

	// When
	sortedKeys := sortLetters(input)

	// Then
	expected := []string{"b:3", "t:3", "x:2"}

	r.Equalf(expected, sortedKeys, "Unexpected output in main()")
}

func TestSortLettersWithSpace(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	input := map[rune]int{rune(' '): 2, rune('x'): 2, rune('b'): 3, rune('t'): 3}

	// When
	sortedKeys := sortLetters(input)

	// Then
	expected := []string{" :2", "b:3", "t:3", "x:2"}

	r.Equalf(expected, sortedKeys, "Unexpected output in main()")
}

func TestSortLettersWithNonAscii(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	input := map[rune]int{rune('®'): 2, rune('‡'): 3, rune('€'): 3}

	// When
	sortedKeys := sortLetters(input)

	// Then
	expected := []string{"®:2", "‡:3", "€:3"}

	r.Equalf(expected, sortedKeys, "Unexpected output in main()")
}
