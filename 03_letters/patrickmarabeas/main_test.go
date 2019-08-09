package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "a:2\nb:1"
	got := buf.String()

	t.Run("Main function", func(t *testing.T) {
		if expected != got {
			t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
		}
	})
}

func TestLetters(t *testing.T) {
	var cases = map[string]struct {
		input    string
		expected map[rune]int
	}{
		"Letters": {
			input:    "abba",
			expected: map[rune]int{'a': 2, 'b': 2},
		},
		"Long": {
			input:    "ffeedcbaaaa",
			expected: map[rune]int{'a': 4, 'b': 1, 'c': 1, 'd': 1, 'e': 2, 'f': 2},
		},
		"Uppercase": {
			input:    "ABBA",
			expected: map[rune]int{'A': 2, 'B': 2},
		},
		"Mixed case": {
			input:    "abBA",
			expected: map[rune]int{'a': 1, 'b': 1, 'A': 1, 'B': 1},
		},
		"Numbers": {
			input:    "1221",
			expected: map[rune]int{'1': 2, '2': 2},
		},
		"emoticons": {
			input:    "✌✌",
			expected: map[rune]int{'✌': 2},
		},
		"empty": {
			input:    "",
			expected: map[rune]int{},
		},
	}

	for name, c := range cases {
		got, expected := letters(c.input), c.expected
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("\nExpected: %q\nGot:      %q", expected, got)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	var cases = map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"Letters": {
			input:    map[rune]int{'b': 2, 'a': 2},
			expected: []string{"a:2", "b:2"},
		},
		"Long": {
			input:    map[rune]int{'a': 4, 'b': 1, 'c': 1, 'd': 1, 'e': 2, 'f': 2},
			expected: []string{"a:4", "b:1", "c:1", "d:1", "e:2", "f:2"},
		},
		"Uppercase": {
			input:    map[rune]int{'A': 2, 'B': 2},
			expected: []string{"A:2", "B:2"},
		},
		"Mixed case": {
			input:    map[rune]int{'a': 1, 'b': 1, 'A': 1, 'B': 1},
			expected: []string{"A:1", "B:1", "a:1", "b:1"},
		},
		"Numbers": {
			input:    map[rune]int{'2': 2, '1': 2},
			expected: []string{"1:2", "2:2"},
		},
		"emoticons": {
			input:    map[rune]int{'✌': 2},
			expected: []string{"✌:2"},
		},
		"empty": {
			input:    map[rune]int{},
			expected: []string{},
		},
	}

	for name, c := range cases {
		got, expected := sortLetters(c.input), c.expected
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("\nExpected: %q\nGot:      %q", expected, got)
			}
		})
	}
}
