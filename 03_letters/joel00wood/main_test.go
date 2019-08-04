package main

import (
	"reflect"
	"testing"

	output "github.com/joel00wood/test-helpers/capture"
)

func TestMain(t *testing.T) {
	expected := "a:2\nb:1\n"
	actual := output.CaptureOutput(func() { main() })
	if expected != actual {
		t.Errorf("Unexpected result in main(), expected=%q, got=%q", expected, actual)
	}
}

func TestLetters(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expected    map[rune]int
		expectedStr []string
	}{
		"Standard input": {
			"aba",
			map[rune]int{97: 2, 98: 1},
			[]string{"a:2", "b:1"},
		},
		"Empty input": {
			"",
			map[rune]int{},
			[]string{},
		},
		"Not so standard input": {
			"The quick brown fox jumped over the two lazy dogs",
			map[rune]int{32: 9, 84: 1, 97: 1, 98: 1, 99: 1, 100: 2, 101: 4, 102: 1, 103: 1,
				104: 2, 105: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1, 111: 5, 112: 1, 113: 1,
				114: 2, 115: 1, 116: 2, 117: 2, 118: 1, 119: 2, 120: 1, 121: 1, 122: 1},
			[]string{" :9", "T:1", "a:1", "b:1", "c:1", "d:2", "e:4", "f:1", "g:1", "h:2", "i:1",
				"j:1", "k:1", "l:1", "m:1", "n:1", "o:5", "p:1", "q:1", "r:2", "s:1", "t:2", "u:2", "v:1",
				"w:2", "x:1", "y:1", "z:1"},
		},
		"Even less standard input": {
			"Th3 qu1ck br0wn faux jump3d ov4 t#3 2 l4zY dawgs?!",
			map[rune]int{32: 9, 33: 1, 35: 1, 48: 1, 49: 1, 50: 1, 51: 3, 52: 2, 63: 1, 84: 1, 89: 1,
				97: 2, 98: 1, 99: 1, 100: 2, 102: 1, 103: 1, 104: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1,
				111: 1, 112: 1, 113: 1, 114: 1, 115: 1, 116: 1, 117: 3, 118: 1, 119: 2, 120: 1, 122: 1},
			[]string{" :9", "!:1", "#:1", "0:1", "1:1", "2:1", "3:3", "4:2", "?:1", "T:1",
				"Y:1", "a:2", "b:1", "c:1", "d:2", "f:1", "g:1", "h:1", "j:1", "k:1", "l:1", "m:1",
				"n:1", "o:1", "p:1", "q:1", "r:1", "s:1", "t:1", "u:3", "v:1", "w:2", "x:1", "z:1"},
		},
	}
	for name, test := range testCases {
		actual := letters(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Unexpected result for letters(%q), expected=%v, got=%v",
				name, test.expected, actual)
		}
	}
}

func TestSortLetters(t *testing.T) {
	testCases := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"Sorted input": {
			map[rune]int{97: 2, 98: 1},
			[]string{"a:2", "b:1"},
		},
		"Empty input": {
			map[rune]int{},
			[]string{},
		},
		"Not so sorted input": {
			map[rune]int{32: 9, 84: 1, 97: 1, 98: 1, 99: 1, 100: 2, 101: 4, 102: 1, 103: 1,
				114: 2, 115: 1, 116: 2, 117: 2, 118: 1, 119: 2, 120: 1, 121: 1, 122: 1, 104: 2,
				105: 1, 106: 1, 107: 1, 108: 1, 109: 1, 110: 1, 111: 5, 112: 1, 113: 1},
			[]string{" :9", "T:1", "a:1", "b:1", "c:1", "d:2", "e:4", "f:1", "g:1", "h:2", "i:1",
				"j:1", "k:1", "l:1", "m:1", "n:1", "o:5", "p:1", "q:1", "r:2", "s:1", "t:2", "u:2", "v:1",
				"w:2", "x:1", "y:1", "z:1"},
		},
		"Even less sorted input": {
			map[rune]int{110: 1, 111: 1, 112: 1, 113: 1, 114: 1, 115: 1, 116: 1, 117: 3, 118: 1, 119: 2,
				120: 1, 97: 2, 98: 1, 99: 1, 100: 2, 102: 1, 103: 1, 104: 1, 106: 1, 107: 1, 108: 1, 109: 1,
				122: 1, 32: 9, 33: 1, 35: 1, 48: 1, 49: 1, 50: 1, 51: 3, 52: 2, 63: 1, 84: 1, 89: 1},
			[]string{" :9", "!:1", "#:1", "0:1", "1:1", "2:1", "3:3", "4:2", "?:1", "T:1",
				"Y:1", "a:2", "b:1", "c:1", "d:2", "f:1", "g:1", "h:1", "j:1", "k:1", "l:1", "m:1",
				"n:1", "o:1", "p:1", "q:1", "r:1", "s:1", "t:1", "u:3", "v:1", "w:2", "x:1", "z:1"},
		},
	}
	for name, test := range testCases {
		actual := sortLetters(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Unexpected result for sortLetters(%q), expected=%v, got=%v",
				name, test.expected, actual)
		}
	}
}
