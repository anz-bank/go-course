package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "[a11y K8s abc]\n"
	actual := buf.String()

	assert.Equal(t, expected, actual)
}

func TestConvertStringToNumeronym(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Empty input", input: "", expected: ""},
		{name: "Two letters word", input: "to", expected: "to"},
		{name: "Three letters word", input: "win", expected: "win"},
		{name: "All lower cases", input: "accessibility", expected: "a11y"},
		{name: "Word begins with capital letter", input: "Kubernetes", expected: "K8s"},
		{name: "Two words with one space", input: "Web Development", expected: "W12t"},
		{name: "Word with non alphanumeric", input: "w$$t", expected: "w$$t"},
	}

	for _, test := range tests {
		actual := convertStringToNumeronym(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
