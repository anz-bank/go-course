package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())

	assert.Equalf(t, expected, actual, "Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)

}

func TestNumeronymOutput(t *testing.T) {
	testCases := map[string]struct {
		input    []rune
		expected string
	}{
		"single string long": {input: []rune("abcdef"),
			expected: "a4f"},
		"string containing non-alphabet": {input: []rune("abs12$ab"),
			expected: "a6b"},
		"non-alphabet only": {input: []rune("$123124"),
			expected: "$54"},
	}

	for name, test := range testCases {
		test := test
		name := name

		t.Run(name, func(t *testing.T) {
			actual := numeronym(test.input)
			assert.Equalf(t, actual, test.expected, "Unexpected output for %v\nexpected: %v,\nactual: %v",
				name, test.expected, actual)
		})
	}
}

func TestNumeronymsOutput(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		expected []string
	}{
		"single string short": {input: []string{"abc"},
			expected: []string{"abc"}},
		"single string long": {input: []string{"abcdef"},
			expected: []string{"a4f"}},
		"multiple strings": {input: []string{"international", "help", "welcome"},
			expected: []string{"i11l", "h2p", "w5e"}},
		"empty string": {input: []string{""},
			expected: []string{""}},
		"string containing non-alphabet": {input: []string{"abs12ab"},
			expected: []string{"a5b"}},
		"non-alphabet only": {input: []string{"123124"},
			expected: []string{"144"}},
		"single emoji": {input: []string{"😀"},
			expected: []string{"😀"}},
		"alphabet + emoji": {input: []string{"a😀bc"},
			expected: []string{"a2c"}},
	}

	for name, test := range testCases {
		test := test
		name := name

		t.Run(name, func(t *testing.T) {
			actual := numeronyms(test.input...)
			assert.Equalf(t, actual, test.expected, "Unexpected output for %v\nexpected: %v,\nactual: %v",
				name, test.expected, actual)
		})
	}
}
