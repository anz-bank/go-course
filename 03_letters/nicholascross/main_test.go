package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main() %v", actual)
	}
}

func TestSortLetters(t *testing.T) {
	var testCases = map[string]struct {
		input map[rune]int
		want  []string
	}{
		"correct counts": {
			input: map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5},
			want:  []string{"a:1", "b:2", "c:3", "d:4", "e:5"},
		},
		"sorted counts": {
			input: map[rune]int{'c': 3, 'd': 4, 'e': 5, 'a': 1, 'b': 2},
			want:  []string{"a:1", "b:2", "c:3", "d:4", "e:5"},
		},
		"empty": {input: map[rune]int{}, want: []string{}},
	}

	for name, test := range testCases {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, sortLetters(input), want, "%v was not sorted", input)
		})
	}
}

func TestLetters(t *testing.T) {
	var testCases = map[string]struct {
		input string
		want  map[rune]int
	}{
		"correct counts": {
			input: "abbcccddddeeeee",
			want:  map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5},
		},
		"unordered counts": {
			input: "ebdbcaddecdecee",
			want:  map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5},
		},
		"special characters": {
			input: "无马e😊e马👍👍a👌😒",
			want:  map[rune]int{'a': 1, '马': 2, '👍': 2, '👌': 1, 'e': 2, '无': 1, '😊': 1, '😒': 1},
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
