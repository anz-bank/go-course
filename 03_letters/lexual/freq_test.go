package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{
			"abcbcc",
			map[rune]int{
				'a': 1,
				'b': 2,
				'c': 3,
			},
		},
		{
			`abcbcc⌘`,
			map[rune]int{
				'a': 1,
				'b': 2,
				'c': 3,
				'⌘': 1,
			},
		},
		{
			`abcbcc丂`,
			map[rune]int{
				'a': 1,
				'b': 2,
				'c': 3,
				'丂': 1,
			},
		},
		{
			"",
			map[rune]int{},
		},
	}
	for _, test := range tests {
		actual := letters(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Fail!\nActual: %v\nExpected: %v", actual, test.expected)
		}
	}
}

func TestSortLetters(t *testing.T) {
	tests := []struct {
		input    map[rune]int
		expected []string
	}{
		{
			map[rune]int{
				'a': 2,
				'b': 1,
				'c': 3,
			},
			[]string{"a:2", "b:1", "c:3"},
		},
		{
			map[rune]int{},
			[]string{},
		},
	}
	for _, test := range tests {
		actual := sortLetters(test.input)
		if len(test.expected) != len(actual) {
			t.Errorf("expected and actual different lengths")
		}
		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Errorf("actual and expected don't match")
				t.Errorf("\nActual: %s\nExpected: %s", actual, test.expected)
			}
		}
	}
}
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
		t.Errorf("\nActual: %s\nExpected: %s", actual, expected)
	}
}
