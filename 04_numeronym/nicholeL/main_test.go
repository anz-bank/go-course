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
	if !assert.Equal(t, strconv.Quote("[a11y K8s abc]"), strconv.Quote(buf.String())) {
		return
	}
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("😀😀😀🤓🤓🤓哈哈哈", "asd", "accessibility")
	expected := []string{"😀7哈", "asd", "a11y"}
	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected[i], v)
		}
	}
}

func TestNumeronym(t *testing.T) {
	testCases := []struct {
		name     string
		input    []rune
		expected string
	}{
		{"Empty string", []rune{' '}, " "},
		{"length of string is one", []rune{'s'}, "s"},
		{"length of string is two", []rune{'s', 's'}, "ss"},
		{"length of string is three", []rune{'a', 's', 'd'}, "asd"},
		{"length of string more than 3", []rune{'d', 's', 'a', 'd', 'a', 's', 'd', 'a', 's', 'f', 'f',
			'j', 'k', 'j', 'n', 'n', 'v', 'c', 'c', 'd', 'f'}, "d19f"},
		{"string contain emojis and special text", []rune{'😀', '😀', '😀', '🤓', '🤓', '🤓',
			'哈', '哈', '哈'}, "😀7哈"},
	}

	for _, testCase := range testCases {
		expected := testCase.expected
		actual := numeronym(testCase.input)
		if !assert.Equal(t, testCase.expected, actual) {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
		}
	}
}
