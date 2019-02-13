package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLettersOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := " :3\nS:1\nT:1\na:1\ng:1\nh:1\ni:3\nn:1\nr:1\ns:2\nt:1\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestLettersLength(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := letters("aaa")

	// Then
	expected := 1
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestLettersOccurance(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := letters("aaa")

	// Then
	expected := 3
	actual := m['a']
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestSortedLetters(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := letters("abaa")
	k := sortedLetters(m)

	// Then
	expected := "a:3"
	actual := k[0]
	r.Equalf(expected, actual, "Unexpected output in main()")
}
