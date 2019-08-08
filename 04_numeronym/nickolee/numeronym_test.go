package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[a11y K8s abc]\n"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
func TestNumeronymsFunc(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected []string
	}

	tests := []test{
		{name: "three strings", input: []string{"horse", "tomato", "potato"}, expected: []string{"h3e", "t4o", "p4o"}},
		{name: "eight mixed strings", input: []string{"tea", "go", "toyota", "ferrari", "a", "Selenium", "aproN", "candy"},
			expected: []string{"tea", "go", "t4a", "f5i", "a", "S6m", "a3N", "c3y"}},
		{name: "empty strings", input: []string{"", "", ""}, expected: []string{"", "", ""}},
	}

	for _, testCase := range tests {
		actual := numeronyms(testCase.input...)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestNumeronyms(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}

	tests := []test{
		{name: "big string", input: "deoxyribonucleicacid", expected: "d18d"},
		{name: "four char string", input: "food", expected: "f2d"},
		{name: "empty string", input: "", expected: ""},
		{name: "three char string", input: "sky", expected: "sky"},
		{name: "two char string", input: "be", expected: "be"},
		{name: "strings with non alphabets", input: "spider-man", expected: "s8n"},
		{name: "starting upper case", input: "Batman", expected: "B4n"},
		{name: "strings with emojis", input: "a👍👍z", expected: "a2z"},
		{name: "strings with emojis", input: "👍👍👍", expected: "👍👍👍"},
		{name: "unicode string with empty space", input: "👍👍👍$€₡ 👍", expected: "👍6👍"},
		{name: "unicode string of numbers", input: "123456789", expected: "179"},
	}

	for _, testCase := range tests {
		actual := numeronym(testCase.input)
		if actual != testCase.expected {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
