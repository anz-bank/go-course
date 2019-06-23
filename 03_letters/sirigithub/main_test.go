package main

import (
	"reflect"
	"testing"
)

/**
Tests to write
**/
func TestLetters(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected map[rune]int
	}{
		{
			"Simple string",
			"accdceee",
			map[rune]int{'a': 1, 'c': 3, 'd': 1, 'e': 3},
		},
		{
			"String with spaces",
			"   223  ",
			map[rune]int{' ': 5, '2': 2, '3': 1},
		},
		{
			"Empty String",
			"",
			map[rune]int{},
		},
		{
			"Empty String",
			"",
			map[rune]int{},
		},
		{
			"German Umlauts",
			"äöß€’üüüöäßß",
			map[rune]int{'ß': 3, 'ä': 2, 'ö': 2, 'ü': 3, '’': 1, '€': 1},
		},
	}

	for _, test := range tests {
		actual := letters(test.input)
		if !(reflect.DeepEqual(actual, test.expected)) {
			t.Errorf("Unexpected output in main()\nexpected: %d\nactual: %d", test.expected, actual)
		}
	}
}
