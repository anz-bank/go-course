package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {

	var tests = map[string]struct {
		arg  string
		want map[rune]int
	}{
		"empty string": {
			arg:  "",
			want: map[rune]int{}},
		"1 element only string": {
			arg:  "a",
			want: map[rune]int{'a': 1}},
		"1 Unicode element only string": {
			arg:  "🤪",
			want: map[rune]int{'🤪': 1}},
		"simple string": {
			arg:  "abacadaba",
			want: map[rune]int{'a': 5, 'b': 2, 'c': 1, 'd': 1}},
		"crazy Unicode chars": {
			arg:  "😀﷽﷽🤪😀c﷽😀d﷽😀🤪😀",
			want: map[rune]int{'😀': 5, '🤪': 2, 'c': 1, 'd': 1, '﷽': 4}},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got := letters(test.arg)
			assert.Equal(t, got, test.want)
		})
	}
}

func TestSortLetters(t *testing.T) {

	var tests = map[string]struct {
		arg  map[rune]int
		want []string
	}{
		"1 element only string": {
			arg:  map[rune]int{'a': 1},
			want: []string{"a:1"}},
		"1 Unicode element only string": {
			arg:  map[rune]int{'🤪': 1},
			want: []string{"🤪:1"}},
		"simple string": {
			arg:  map[rune]int{'b': 2, 'a': 5, 'd': 1, 'c': 1},
			want: []string{"a:5", "b:2", "c:1", "d:1"}},
		"crazy Unicode chars": {
			arg:  map[rune]int{'🤪': 2, 'c': 1, 'd': 1, '﷽': 4},
			want: []string{"c:1", "d:1", "﷽:4", "🤪:2"}},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got := sortLetters(test.arg)
			assert.Equal(t, got, test.want)
		})
	}
}

func TestMain(t *testing.T) {
	want := "a:2\nb:1\n"
	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, got, want)
	})
}
