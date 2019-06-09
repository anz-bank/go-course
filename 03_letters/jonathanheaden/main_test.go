package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	// Given
	a := assert.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	a.Equalf(expected, actual, "Unexpected output in main()")
}

func TestLetters(t *testing.T) {
	var testCases = map[string]struct {
		input string
		want  map[rune]int
	}{
		"correct counts": {
			input: "abracadabra",
			want:  map[rune]int{'a': 5, 'b': 2, 'c': 1, 'd': 1, 'r': 2},
		},
		"mixed case": {
			input: "HHhhEeelellLLoo0",
			want:  map[rune]int{'0': 1, 'E': 1, 'H': 2, 'L': 2, 'e': 3, 'h': 2, 'l': 3, 'o': 2},
		},
		"empty": {input: "", want: map[rune]int{}},
	}

	for name, test := range testCases {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, letters(input), want, "%v was not correct", input)
		})
	}
}
