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

	expected := strconv.Quote(`[a11y K8s abc]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("testString", "www", "W3school")
	expected := []string{"t8g", "www", "W6l"}
	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected[i], v)
		}
	}
}

func TestGetNumeronyms(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Empty string": {
			input:    "",
			expected: "",
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
			actual := getNumeronym(input)
			assert.Equalf(t, expected, actual, "Unexpected output")
		})
	}
}
