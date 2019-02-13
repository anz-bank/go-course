package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetters(t *testing.T) {
	// Given
	r := require.New(t)
	// Then

	expected := map[rune]int{97: 2, 98: 1}
	actual := letters("aba")
	r.Equal(expected, actual)

}

func TestSortedLetters(t *testing.T) {
	// Given
	r := require.New(t)
	// Then

	expected := []string{"a : 2", "b : 1"}
	actual := sortLetters(map[rune]int{97: 2, 98: 1})
	r.Equal(expected, actual)

}
func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[a : 2 b : 1]\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
