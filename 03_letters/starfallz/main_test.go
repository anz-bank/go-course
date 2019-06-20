package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of letters(aba) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote(`a:2
b:1`)
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestLettersFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    map[rune]int
	}{
		{"Test letters to return sorted map of letters based on rune # with count", "abaabazzzzzz",
			map[rune]int{97: 4, 98: 2, 122: 6}},
		{"Test letters to handle uppercase letters", "aAAaA", map[rune]int{65: 3, 97: 2}},
		{"Test letters to handle non-alphabetical characters", "    !!!*", map[rune]int{32: 4, 33: 3, 42: 1}},
		{"Test letters to return empty map on empty string input", "", map[rune]int{}},
	}

	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected
		t.Run(testCase.description, func(t *testing.T) {
			actual := letters(input)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Unexpected output, expected: %d, actual: %d", expected, actual)
			}
		})
	}
}

func TestSortLettersFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       map[rune]int
		expected    []string
	}{
		{"Test sort letters using letters() as input and return count", letters("aba"), []string{"a:2", "b:1"}},
		{"Test sort letters with uppercase using letters() as input and return count", letters("Aba"),
			[]string{"A:1", "a:1", "b:1"}},
		{"Test sort letters with uppercase and return count", map[rune]int{97: 2, 65: 3}, []string{"A:3", "a:2"}},
		{"Test sort letters with empty slice", map[rune]int{}, []string{}},
		{"Test nil input to output nil", nil, []string{}},
	}

	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected
		t.Run(testCase.description, func(t *testing.T) {
			actual := sortLetters(input)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
			}
		})
	}
}

func TestSortRuneFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       []rune
		expected    []rune
	}{
		{"Test sort unsorted rune slice to sorted rune slice", []rune{97, 82}, []rune{82, 97}},
		{"Test sort negative rune", []rune{-97, 82}, []rune{-97, 82}},
		{"Test sort empty slice", []rune{}, []rune{}},
		{"Test nil input to output nil", nil, nil},
	}

	for _, testCase := range testCases {
		actual := testCase.input
		expected := testCase.expected
		t.Run(testCase.description, func(t *testing.T) {
			sortRune(actual)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Unexpected output, expected: %d, actual: %d", expected, actual)
			}
		})
	}
}
