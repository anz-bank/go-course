package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	s    string
	lmap map[rune]int
	ss   []string
}{

	{"pwopoip@oiopoi@",
		map[rune]int{'o': 5, 'p': 4, 'i': 3, '@': 2, 'w': 1},
		[]string{"@:2", "i:3", "o:5", "p:4", "w:1"}},
	{"‣£ê‣££~ 3 3 3",
		map[rune]int{'3': 3, '£': 3, ' ': 3, '‣': 2, 'ê': 1, '~': 1},
		[]string{" :3", "3:3", "~:1", "£:3", "ê:1", "‣:2"}},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestLettersFrequency(t *testing.T) {
	r := require.New(t)
	for _, v := range tests {
		expectedMap := letters(v.s)
		r.Equal(v.lmap, expectedMap)
		expectedString := sortLetters(v.lmap)
		r.EqualValues(v.ss, expectedString)
	}
}
