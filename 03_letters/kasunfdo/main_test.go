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

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestLetters(t *testing.T) {
	testData := "Two Driven Jocks Help Fax My Big Quiz!"

	expected := map[rune]int{
		'i': 3, 'k': 1, 'l': 1, 'n': 1, 'o': 2, 'p': 1, 'r': 1, 's': 1, 'H': 1, 'J': 1, 'M': 1, 'Q': 1, 'u': 1, 'v': 1,
		'w': 1, 'x': 1, 'y': 1, ' ': 7, '!': 1, 'B': 1, 'D': 1, 'F': 1, 'T': 1, 'a': 1, 'c': 1, 'e': 2, 'g': 1, 'z': 1,
	}

	actual := letters(testData)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestSortLetters(t *testing.T) {
	testData := "Two Driven Jocks Help Fax My Big Quiz!"

	expected := []string{
		" :7", "!:1", "B:1", "D:1", "F:1", "H:1", "J:1", "M:1", "Q:1", "T:1", "a:1", "c:1", "e:2", "g:1",
		"i:3", "k:1", "l:1", "n:1", "o:2", "p:1", "r:1", "s:1", "u:1", "v:1", "w:1", "x:1", "y:1", "z:1"}

	actual := sortLetters(letters(testData))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestRandomStrWithSymbols(t *testing.T) {

}
