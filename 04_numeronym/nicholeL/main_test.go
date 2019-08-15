package main

import (
	"bytes"

	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("[a11y K8s abc]")
	actual := strconv.Quote(buf.String())
	if actual != expected {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestNumberShorten(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"Happy case":                {"accessibility", "a11y"},
		"length of string is three": {"asd", "asd"},
		"UnHappy case":              {"", ""},
	}
	for key, testCase := range testCases {
		expected := testCase.expected
		actual := numberShorten(testCase.input)
		t.Run(key, func(t *testing.T) {
			if expected != actual {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
			}
		})
	}
}

func TestNumeronyms(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"length of string is one", "s", "s"},
		{"length of string is two", "ss", "ss"},
		{"length of string is three", "asd", "asd"},
		{"length of string more than 3", "dsadasdasffjkjnnvccdf", "d19f"},
	}

	for _, testCase := range testCases {
		expected := testCase.expected
		actual := numberShorten(testCase.input)
		if !assert.Equal(t, testCase.expected, actual) {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
		}
	}
}
