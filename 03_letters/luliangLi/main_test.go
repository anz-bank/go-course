package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestSortLetters(t *testing.T) {

	testCases := map[string]struct {
		input string
		want  []string
	}{
		"normal": {input: "aba", want: []string{"a:2", "b:1"}},
		"blank":  {input: "", want: []string{}},
		"long":   {input: "aabbcccdddd", want: []string{"d:4", "c:3", "a:2", "b:2"}},
		"single": {input: "aaaa", want: []string{"a:4"}},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		test := test

		t.Run(name, func(t *testing.T) {
			actual := sortLetters(letters(test.input))

			if !reflect.DeepEqual(test.want, actual) {
				t.Errorf("running : %v, expected %v, got %v", test.input, test.want, actual)
			}
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("running main, expected %v, got %v", expected, actual)
	}
}
