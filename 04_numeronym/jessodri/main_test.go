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
	assert.Equal(t, actual, expected)
}

func TestNumeronyms(t *testing.T) {
	testCases := []struct {
		description string
		word        []string
		expected    []string
	}{
		{
			"happy path - one string",
			[]string{"onomatopoeia"},
			[]string{"o10a"},
		},
		{
			"happy path - multiple strings",
			[]string{"onomatopoeia", "K8bernetes", "dog", "chicken", "four"},
			[]string{"o10a", "K8s", "dog", "c5n", "f2r"},
		},
		{
			"3 letter word",
			[]string{"cat"},
			[]string{"cat"},
		},
		{
			"empty",
			[]string{""},
			[]string{""},
		},
		{
			"multibyte string",
			[]string{"世界世界", "世界世界"},
			[]string{"世2界", "世2界"},
		},
	}

	for _, tc := range testCases {
		expected := tc.expected
		actual := numeronyms(tc.word...)
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, actual, expected, "got %v but wanted %v", actual, expected)
		})
	}
}
