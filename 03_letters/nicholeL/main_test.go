package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	excepted := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	if excepted != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
	}

}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input    string
		excepted map[rune]int
	}{
		"one":   {"aba", map[rune]int{97: 2, 98: 1}},
		"two":   {"abadddcca", map[rune]int{97: 3, 98: 1, 99: 2, 100: 3}},
		"three": {"abadddcc311a", map[rune]int{'1': 2, '3': 1, 97: 3, 98: 1, 99: 2, 100: 3}},
		"four":  {"ccccccabadddcca", map[rune]int{97: 3, 98: 1, 99: 8, 100: 3}},
		"five":  {"", map[rune]int{}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := letters(test.input)
			if !reflect.DeepEqual(test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input    map[rune]int
		excepted []string
	}{
		"one": {map[rune]int{97: 2, 98: 1}, []string{"a:2", "b:1"}},
		"two": {map[rune]int{97: 3, 98: 1, 99: 2, 100: 3}, []string{"a:3", "b:1", "c:2", "d:3"}},
		"three": {map[rune]int{1: 2, 3: 1, 97: 3, 98: 1, 99: 2, 100: 3}, []string{"\x01:2", "\x03:1",
			"a:3", "b:1", "c:2", "d:3"}},
		"four": {map[rune]int{97: 3, 98: 1, 99: 8, 100: 3, 67: 1}, []string{"C:1", "a:3", "b:1",
			"c:8", "d:3"}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := sortLetters(test.input)
			if !reflect.DeepEqual(test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})
	}
}
