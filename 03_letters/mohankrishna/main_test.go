package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	inputString   string
	lettersMap    map[rune]int
	sortedStrings []string
}{
	{"", map[rune]int{}, []string{}},
	{"abdsd3$asda$asasdd$sadas",
		map[rune]int{'a': 7, 'b': 1, 'd': 6, 's': 6, '3': 1, '$': 3},
		[]string{"$:3", "3:1", "a:7", "b:1", "d:6", "s:6"}},
	{"!&^#*&+@#*&@+#$#&!*$!$ !^+ &$!*&$^ +$&^!#$ &!$!^$&!# $",
		map[rune]int{'!': 9, '&': 9, '^': 5, '#': 6, '*': 4, '+': 4, '@': 2, '$': 10, ' ': 5},
		[]string{" :5", "!:9", "#:6", "$:10", "&:9", "*:4", "+:4", "@:2", "^:5"}},
	{"🦋😱😈😎🦖🤠😎😈🦖😎🐉🤠🦖🦋🦋🦋😈😎🐉🦖🤠😈🐉😎🤠😈🐉🤠😱🦋",
		map[rune]int{'🦋': 5, '😱': 2, '😈': 5, '😎': 5, '🦖': 4, '🤠': 5, '🐉': 4},
		[]string{"🐉:4", "😈:5", "😎:5", "😱:2", "🤠:5", "🦋:5", "🦖:4"}}}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `a:2
b:1
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestLettersMapAndSort(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		outMap := letters(tt.inputString)
		r.Equal(tt.lettersMap, outMap)
		outStrings := sortLetters(tt.lettersMap)
		r.EqualValues(tt.sortedStrings, outStrings)
	}
}
