package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "a:2\nb:1\n", buf.String())
}

func TestLetters(t *testing.T) {
	assert.Equal(t, map[rune]int{' ': 2}, letters("  "))
	assert.Equal(t, map[rune]int{'a': 4, 'b': 2}, letters("abaaba"))
	assert.Equal(t, map[rune]int{'m': 1, 'i': 4, 's': 4, 'p': 2}, letters("mississippi"))
	assert.Equal(t,
		map[rune]int{' ': 3, 'c': 2, 'h': 2, 'i': 4, 'k': 2, 'b': 2, 'u': 2, 'm': 2},
		letters("chiki chiki bum bum"))
}

func TestLettersUnicode(t *testing.T) {
	assert.Equal(t, map[rune]int{'\u2039': 2, '😀': 2, '😇': 1, '😎': 1}, letters("\u2039\u2039😀😀😇😎"))
}

func TestSortLetters(t *testing.T) {
	assert.Equal(t, []string{}, sortLetters(map[rune]int{}))
	assert.Equal(t, []string{" :2"}, sortLetters(map[rune]int{' ': 2}))
	assert.Equal(t, []string{"a:2", "b:4"}, sortLetters(map[rune]int{'b': 4, 'a': 2}))
	assert.Equal(t,
		[]string{"i:4", "m:1", "p:2", "s:4"},
		sortLetters(map[rune]int{'m': 1, 'i': 4, 's': 4, 'p': 2}))
	assert.Equal(t,
		[]string{"a:4", "o:4", "q:1", "v:2"},
		sortLetters(map[rune]int{'q': 1, 'a': 4, 'o': 4, 'v': 2}))
}

func TestSortLettersUnicode(t *testing.T) {
	assert.Equal(t,
		[]string{"\u2020:5", "\u2025:3", "\u2040:1", "😀:1", "😙:2", "🧐:2"},
		sortLetters(map[rune]int{'\u2040': 1, '🧐': 2, '😙': 2, '😀': 1, '\u2025': 3, '\u2020': 5}))
}
