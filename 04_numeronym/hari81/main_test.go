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

	expected := strconv.Quote("[r17g j8t a3i]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("reactjsprogramming", "javascript", "anzui")
	expected := []string{"r16g", "j8t", "a3i"}
	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected[i], v)
		}
	}
}

func TestShorten(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"anzui":              {input: "anzui", expected: "a3i"},
		"javascript":         {input: "javascript", expected: "j8t"},
		"reactjsprogramming": {input: "reactjsprogramming", expected: "r16g"},
	}
	for name, test := range testCases {
		input, want := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			actual := shorten(input)
			assert.Equalf(t, want, actual, "For input: %v expected: %v got %v", input, want, actual)
		})
	}
}
