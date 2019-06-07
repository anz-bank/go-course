package main

import (
	"bytes"
	"fmt"
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

func TestSortLetters(t *testing.T) {
	var bufletter bytes.Buffer
	outletter = &bufletter

	type test struct {
		input    string
		expected []string
		actual   []string
	}

	tests := []test{
		{input: "abac", expected: []string{"a:2", "b:1", "c:1"}, actual: sortLetters(letters("abac"))},
		{input: "aAdbbDA", expected: []string{"A:2", "D:1", "a:1", "b:2", "d:1"}, actual: sortLetters(letters("aAdbbDA"))},
		{input: "aAAADDDbac", expected: []string{"A:3", "D:3", "a:2", "b:1", "c:1"},
			actual: sortLetters(letters("aAAADDDbac"))},
		{input: "", expected: []string{}, actual: sortLetters(letters(""))},
	}

	for _, tc := range tests {
		if len(tc.expected) > 0 && len(tc.actual) > 0 {
			if !reflect.DeepEqual(tc.expected, tc.actual) {
				fmt.Printf("%v  %T", tc.expected, tc.expected)
				fmt.Printf("%v  %T", tc.actual, tc.actual)
				t.Errorf(" input:%v, expected: %v, got: %v", tc.input, tc.expected, tc.actual)
			}
		}
	}

}
