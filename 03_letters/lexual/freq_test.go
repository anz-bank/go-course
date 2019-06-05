package main

import (
	"bytes"
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
			"",
			map[rune]int{},
		},
	}
	for _, test := range tests {
		actual := letters(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("len of actual and expected differ")
		}
		for k, v := range test.expected {
			if v != actual[k] {
				t.Errorf("Unexpected result")
				t.Errorf("\nActual: %v\nExpected: %v", actual, test.expected)
				for k, v := range actual {
					t.Errorf("Actual %v = %v\n", string(k), v)
				}
				for k, v := range test.expected {
					t.Errorf("Expect %v = %v\n", string(k), v)
				}
			}
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
		// empty map as argument tested
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

				break
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
