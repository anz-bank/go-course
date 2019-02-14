package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in       string
	Expected string
}{
	{"mango", "[m3o]"},
	{"This is my test case", "[T18e]"},
	{"This ", "[T2s]"},
	{" This ", "[T2s]"},
	{" This", "[T2s]"},
	{"", "[]"},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[m3o app]"
	actual := buf.String()

	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronymWithInput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := numeronyms("mango", "apple", "bankanywhere")

	// Then
	expected := []string([]string{"m3o", "a3e", "b10e"})
	r.Equalf(expected, s, "Unexpected output in main()")
}

func TestNumeronymWithSpaceInput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := numeronyms(" mango", "apple ", " bankanywhere ")

	// Then
	expected := []string([]string{"m3o", "a3e", "b10e"})
	r.Equalf(expected, s, "Unexpected output in main()")
}

func TestNumeronymWithTwoInputs(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := numeronyms("mango", "apple")

	// Then
	expected := []string([]string{"m3o", "a3e"})
	r.Equalf(expected, s, "Unexpected output in main()")
}

func TestNumeronymWithOneInput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	s := numeronyms("mango")

	// Then
	expected := []string([]string{"m3o"})
	r.Equalf(expected, s, "Unexpected output in main()")
}

func TestNumeronymWithEmptyInput(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	s := numeronyms()

	// Then
	assert.Equal(0, len(s), "Unexpected output in main()")
}

func TestNumeronyms(t *testing.T) {
	// Given
	assert := assert.New(t)

	for _, v := range tests {
		assert.Equalf(v.Expected, fmt.Sprint(numeronyms(v.in)), "Unexpected output in numeronyms()")
	}
}
