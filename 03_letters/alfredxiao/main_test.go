package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCountsEdgeCases(t *testing.T) {
	assert.Equal(t, map[rune]int{}, letters(""))
	assert.Equal(t, map[rune]int{'a': 1}, letters("a"))
	assert.Equal(t, map[rune]int{'a': 3}, letters("aaa"))
	assert.Equal(t, map[rune]int{' ': 1}, letters(" "))
	assert.Equal(t, map[rune]int{' ': 3}, letters("   "))
}

func TestLetterCountsUnicode(t *testing.T) {
	assert.Equal(t, map[rune]int{'\u2318': 1}, letters("\u2318"))
	assert.Equal(t, map[rune]int{'\u2318': 1, '\u2319': 1}, letters("\u2318\u2319"))
	assert.Equal(t, map[rune]int{'\u2318': 1, '\u2319': 1, 'A': 1}, letters("\u2318A\u2319"))
}

func TestLetterCountsGeneralCases(t *testing.T) {
	assert.Equal(t, map[rune]int{'a': 3, 'b': 2, 'c': 1}, letters("cababa"))
}

func TestSortLettersEdgeCases(t *testing.T) {
	assert.Empty(t, sortLetters(map[rune]int{}))
	assert.Equal(t, []string{"a:1"}, sortLetters(map[rune]int{'a': 1}))
	assert.Equal(t, []string{"\u2318:1"}, sortLetters(map[rune]int{'\u2318': 1}))
}

func TestSortLettersGeneralCases(t *testing.T) {
	assert.Equal(t, []string{"a:2", "c:1"}, sortLetters(map[rune]int{'c': 1, 'a': 2}))
	assert.Equal(t, []string{"\u2318:2", "\u2319:1"}, sortLetters(map[rune]int{'\u2319': 1, '\u2318': 2}))
	assert.Equal(t, []string{"a:2", "\u2318:1"}, sortLetters(map[rune]int{'\u2318': 1, 'a': 2}))
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "a:2\nb:1\n", buf.String())
}
