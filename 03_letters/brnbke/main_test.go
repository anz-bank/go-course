package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

var hanMap = map[rune]int{'⽖': 1, '⽉': 2, '⽏': 3, '⽐': 2, '⽕': 1}
var emojiMap = map[rune]int{'🏦': 1, '🔫': 4, '💰': 1, '🚗': 1, '😬': 1,
	'🚓': 7, '😱': 1, '🚒': 1, '🚑': 1, '😨': 1, '😢': 7, '😰': 1,
	'🗯': 1, '🏛': 1, '⏸': 1, '🔒': 1, '👮': 1}
var latinMap = map[rune]int{' ': 8, 'a': 2, 'b': 1, 'c': 1, 'd': 1, 'e': 2,
	'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1, 'l': 1, 'm': 1,
	'n': 1, 'o': 4, 'p': 1, 'q': 1, 'r': 2, 's': 1, 't': 1, 'u': 2,
	'v': 1, 'w': 1, 'x': 1, 'y': 1, 'z': 1}
var canadianMap = map[rune]int{'ᓁ': 1, 'ᓂ': 1, 'ᓃ': 1, 'ᓄ': 3, 'ᓅ': 1, 'ᓐ': 1,
	'ᓑ': 1, 'ᓒ': 1, 'ᓓ': 1, 'ᓔ': 1, 'ᓱ': 1, 'ᓲ': 1, 'ᓳ': 1, 'ᓴ': 1,
	'ᓵ': 1, 'ᓶ': 1, 'ᔒ': 1, 'ᔓ': 1, 'ᔔ': 1, 'ᔕ': 1, 'ᔦ': 1, 'ᔧ': 2}

func TestLetters(t *testing.T) {
	var testCases = map[string]struct {
		input string
		want  map[rune]int
	}{
		"Empty String": {
			input: "",
			want:  map[rune]int{}},
		"Han Script": {
			input: "⽉⽐⽏⽐⽏⽕⽉⽏⽖",
			want:  hanMap},
		"Robbing a bank": {
			input: "🏦🔫🗯💰🚗😬🚓🚓🚓🚓😱🔫🔫🔫🚓🚓🚒🚑😨😢😰😢😢🚓🏛😢😢😢😢⏸🔒👮",
			want:  emojiMap},
		"Some Latin chars": {
			input: "the quick brown fox jumps over a lazy god",
			want:  latinMap},
		"Canadial Aboriginal": {
			input: "ᓁᓂᓃᓄᓅᓐᓑᓒᓓᓔᓱᓄᓄᓲᓳᓴᓵᔧᓶᔒᔓᔔᔕᔦᔧ",
			want:  canadianMap},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			b := letters(test.input)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("want %v, got %v", test.want, b)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	var testCases = map[string]struct {
		input map[rune]int
		want  []string
	}{
		"Empty Map": {
			input: map[rune]int{},
			want:  []string{}},
		"Han Script": {
			input: hanMap,
			want: []string{
				"⽉:2", "⽏:3", "⽐:2", "⽕:1", "⽖:1",
			}},
		"Robbing a bank": {
			input: emojiMap,
			want: []string{
				"⏸:1", "🏛:1", "🏦:1", "👮:1", "💰:1", "🔒:1", "🔫:4",
				"🗯:1", "😢:7", "😨:1", "😬:1", "😰:1", "😱:1", "🚑:1",
				"🚒:1", "🚓:7", "🚗:1"}},
		"Some Latin chars": {
			input: latinMap,
			want: []string{
				" :8", "a:2", "b:1", "c:1", "d:1", "e:2",
				"f:1", "g:1", "h:1", "i:1", "j:1", "k:1", "l:1", "m:1",
				"n:1", "o:4", "p:1", "q:1", "r:2", "s:1", "t:1", "u:2",
				"v:1", "w:1", "x:1", "y:1", "z:1"}},
		"Canadial Aboriginal": {
			input: canadianMap,
			want: []string{
				"ᓁ:1", "ᓂ:1", "ᓃ:1", "ᓄ:3", "ᓅ:1", "ᓐ:1", "ᓑ:1", "ᓒ:1",
				"ᓓ:1", "ᓔ:1", "ᓱ:1", "ᓲ:1", "ᓳ:1", "ᓴ:1", "ᓵ:1", "ᓶ:1",
				"ᔒ:1", "ᔓ:1", "ᔔ:1", "ᔕ:1", "ᔦ:1", "ᔧ:2"}},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			b := sortLetters(test.input)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("want %v, got %v", test.want, b)
			}
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if want != actual {
		t.Errorf("Unwant output in main()\nwant: %q\nactual: %q", want, actual)
	}
}
