package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSortLettersMain(t *testing.T) {
	var bufletter bytes.Buffer
	outletter = &bufletter

	main()
	expected := `a:2
b:1`
	actual := bufletter.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()" + expected + " " + actual)
	}
}

func TestLetters(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected map[rune]int
	}{"Test1": {"abac", map[rune]int{'a': 2, 'b': 1, 'c': 1}},
		"Test2": {"aAdbbDA", map[rune]int{'A': 2, 'D': 1, 'a': 1, 'b': 2, 'd': 1}},
		"Test3": {"aAAADDDbac", map[rune]int{'A': 3, 'D': 3, 'a': 2, 'b': 1, 'c': 1}},
		"Test4": {"", map[rune]int{}},
	}
	for name, test := range cases {
		actual := letters(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(" test:%v input:%v, expected: %v, got: %v", name, test.input, test.expected, actual)
		}
	}

}

func TestSortLetters(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{"Test1": {"abac", []string{"a:2", "b:1", "c:1"}},
		"Test2": {"aAdbbDA", []string{"A:2", "D:1", "a:1", "b:2", "d:1"}},
		"Test3": {"aAAADDDbac", []string{"A:3", "D:3", "a:2", "b:1", "c:1"}},
		"Test4": {"", []string{}},
	}
	for name, test := range cases {
		actual := sortLetters(letters(test.input))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(" test:%v input:%v, expected: %v, got: %v", name, test.input, test.expected, actual)
		}
	}

}
