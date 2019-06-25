package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestNumeronym(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"accessibility", "a11y"},
		{"Kubernetes", "K8s"},
		{"abc", "abc"},
		{"abcd", "a2d"},
		{"ab", "ab"},
		{"a", "a"},
		{"", ""},
	}
	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			actual := numeronym(test.input)
			if actual != test.expected {
				t.Errorf("input: %s, expected: %s, actual: %s",
					test.input, test.expected, actual)
			}
		})
	}
}

func TestNumeronyms(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"3 examples": {
			[]string{"accessibility", "Kubernetes", "abc"},
			[]string{"a11y", "K8s", "abc"},
		},
		"0 examples": {
			[]string{},
			[]string{},
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := numeronyms(test.input...)
			if len(actual) != len(test.expected) {
				t.Errorf("len(expected): %d, len(actual): %d",
					len(test.expected), len(actual))
			}
			for i := range test.expected {
				if actual[i] != test.expected[i] {
					t.Errorf("input: %s, expected: %#v, actual: %#v",
						test.input, test.expected, actual)
					break
				}
			}
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
		t.Errorf("\nActual: %s\nExpected: %s", actual, expected)
	}
}
