package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[a11y K8s abc]"
	got := buf.String()

	t.Run("Main function", func(t *testing.T) {
		if expected != got {
			t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
		}
	})
}

func TestGenerateNumeronym(t *testing.T) {
	var cases = map[string]struct {
		input    string
		expected string
	}{
		"Lowercase": {
			input:    "lowercase",
			expected: "l7e",
		},
		"Uppercase": {
			input:    "Uppercase",
			expected: "U7e",
		},
		"Three letters": {
			input:    "abc",
			expected: "abc",
		},
		"Two letters": {
			input:    "ab",
			expected: "ab",
		},
		"Single letter": {
			input:    "a",
			expected: "a",
		},
		"": {
			input:    "😚😚😚😚😚😚",
			expected: "😚4😚",
		},
		"Empty": {
			input:    "",
			expected: "",
		},
	}

	for name, c := range cases {
		got, expected := generateNumeronym(c.input), c.expected
		t.Run(name, func(t *testing.T) {
			if got != expected {
				t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
			}
		})
	}
}

func TestNumeronyms(t *testing.T) {
	test := numeronyms("lowercase", "Uppercase")
	expected := []string{"l7e", "U7e"}

	for i, got := range test {
		// Lint error: Using the variable on range scope `got` / `i` in function literal
		i := i
		got := got
		t.Run(fmt.Sprintf("String of index %d", i), func(t *testing.T) {
			if got != expected[i] {
				t.Errorf("\nExpected: %s\nGot:      %s", expected[i], got)
			}
		})
	}
}
