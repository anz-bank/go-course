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

	assert.Equalf(t, expected, actual, "Unexpected output")
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("testString", "www", "W3school")
	expected := []string{"t8g", "www", "W6l"}
	for i, v := range actual {
		assert.Equalf(t, expected[i], v, "Unexpected output")
	}
}

func TestNumeronym(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Empty string": {
			input:    "",
			expected: "",
		},
		"3 letter string": {
			input:    "abc",
			expected: "abc",
		},
		"4 letter string": {
			input:    "abcd",
			expected: "a2d",
		},
		"ASCII string": {
			input:    " Multilingualization",
			expected: "M17n",
		},
		"Rune string": {
			input:    " ➳↪▲✓₷$€₡😀⤆ ",
			expected: "➳8⤆",
		},
	}

	for name, test := range tests {
		input, expected := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			actual := numeronym(input)
			assert.Equalf(t, expected, actual, "Unexpected output")
		})
	}
}
