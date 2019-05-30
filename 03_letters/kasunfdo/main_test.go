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
	type test struct {
		input    string
		expected []string
	}

	tests := []test{
		{input: "Two Driven Jocks Help Fax My Big Quiz!",
			expected: []string{" :7", "!:1", "B:1", "D:1", "F:1", "H:1", "J:1", "M:1", "Q:1",
				"T:1", "a:1", "c:1", "e:2", "g:1", "i:3", "k:1", "l:1", "n:1", "o:2", "p:1",
				"r:1", "s:1", "u:1", "v:1", "w:1", "x:1", "y:1", "z:1"}},
		{input: "菎絝覓滃惟哫쉱췡픳㰞섭箕쒐㼪ν㰗뾬✌煡",
			expected: []string{"ν:1", "✌:1", "㰗:1", "㰞:1", "㼪:1", "哫:1", "惟:1", "滃:1", "煡:1",
				"箕:1", "絝:1", "菎:1", "覓:1", "뾬:1", "섭:1", "쉱:1", "쒐:1", "췡:1", "픳:1"}},
	}

	for _, test := range tests {
		actual := sortLetters(letters(test.input))
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual  : %q", test.expected, actual)
		}
	}
}
