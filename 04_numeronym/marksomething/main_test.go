package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeronym(t *testing.T) {
	testCases := map[string]struct {
		arg      string
		expected string
	}{
		"Basic": {
			arg:      "Hello",
			expected: "H3o",
		},
		"Emoji": {
			arg:      "😀💩💩💩🤡🤡",
			expected: "😀4🤡",
		},
		"Len 1": {
			arg:      "A",
			expected: "A",
		},
		"Len 3": {
			arg:      "Cat",
			expected: "Cat",
		},
		"Empty": {
			arg:      "",
			expected: "",
		},
	}

	for testName, tC := range testCases {
		testCase := tC
		t.Run(testName, func(t *testing.T) {
			actual := numeronym(testCase.arg)
			expected := testCase.expected
			assert.Equal(t, expected, actual)
		})
	}
}

func TestNumeronyms(t *testing.T) {
	testCases := map[string]struct {
		arg      []string
		expected []string
	}{
		"Single": {
			arg:      []string{"Hello"},
			expected: []string{"H3o"},
		},
		"Multiple": {
			arg:      []string{"Hello", "Goodbye"},
			expected: []string{"H3o", "G5e"},
		},
		"Empty": {
			arg:      []string{},
			expected: []string{},
		},
	}

	for testName, tC := range testCases {
		testCase := tC
		t.Run(testName, func(t *testing.T) {
			actual := numeronyms(testCase.arg...)
			expected := testCase.expected
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	actual := buf.String()
	expected := "[a11y K8s abc]\n"
	assert.Equal(t, expected, actual)
}
