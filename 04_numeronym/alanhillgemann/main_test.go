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
	expected := strconv.Quote("[a11y K8s abc]")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
}

func TestNumeronyms(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		expected []string
	}{
		"0 params":     {input: []string{}, expected: []string{}},
		"0 char":       {input: []string{""}, expected: []string{""}},
		"1 char":       {input: []string{"a"}, expected: []string{"a"}},
		"2 chars":      {input: []string{"ab"}, expected: []string{"ab"}},
		"3 chars":      {input: []string{"abc"}, expected: []string{"abc"}},
		"4 chars":      {input: []string{"abcd"}, expected: []string{"a2d"}},
		"> 11 chars":   {input: []string{"abcdefghijkl"}, expected: []string{"a10l"}},
		"utf8 char":    {input: []string{"d€f"}, expected: []string{"d€f"}},
		"multi params": {input: []string{"abcef", "defghi"}, expected: []string{"a3f", "d4i"}},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := numeronyms(test.input...)
			expected := test.expected
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}
