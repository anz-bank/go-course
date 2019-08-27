package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var letterTests = []struct {
	input string
	want  map[rune]int
}{
	{input: "abcdefghijk",
		want: map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1}},
	{input: "894431kdk",
		want: map[rune]int{'8': 1, '9': 1, '4': 2, '3': 1, '1': 1, 'k': 2, 'd': 1}},
	{input: "好好学习天天向上",
		want: map[rune]int{'好': 2, '学': 1, '习': 1, '天': 2, '向': 1, '上': 1}},
	{input: "🏦🔫🗯💰😬🚓🚓🚓🚓",
		want: map[rune]int{'🏦': 1, '🔫': 1, '🗯': 1, '💰': 1, '😬': 1, '🚓': 4}},
}

var sortletterTests = []struct {
	input map[rune]int
	want  []string
}{
	{input: map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1},
		want: []string{"a:1", "b:1", "c:1", "d:1", "e:1", "f:1", "g:1", "h:1", "i:1", "j:1", "k:1"}},
	{input: map[rune]int{'8': 1, '9': 1, '4': 2, '3': 1, '1': 1, 'k': 2, 'd': 1},
		want: []string{"1:1", "3:1", "4:2", "8:1", "9:1", "d:1", "k:2"}},
	{input: map[rune]int{'好': 2, '学': 1, '习': 1, '天': 2, '向': 1, '上': 1},
		want: []string{"上:1", "习:1", "向:1", "天:2", "好:2", "学:1"}},
	{input: map[rune]int{'🏦': 1, '🔫': 1, '🗯': 1, '💰': 1, '😬': 1, '🚓': 4},
		want: []string{"🏦:1", "💰:1", "🔫:1", "🗯:1", "😬:1", "🚓:4"}},
}

func TestLetters(t *testing.T) {
	for _, test := range letterTests {
		actual := letters(test.input)
		assert.Equal(t, test.want, actual)
	}
}

func TestSortLetters(t *testing.T) {
	for _, test := range sortletterTests {
		actual := sortLetters(test.input)
		assert.Equal(t, test.want, actual)
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
