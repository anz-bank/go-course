package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		numInput       int
		expectedOutput string
	}{
		{numInput: 7, expectedOutput: "1 1 2 3 5 8 13 "},
		{numInput: -7, expectedOutput: "1 -1 2 -3 5 -8 13 "},
		{numInput: 10, expectedOutput: "1 1 2 3 5 8 13 21 34 55 "},
		{numInput: -10, expectedOutput: "1 -1 2 -3 5 -8 13 -21 34 -55 "},
		{numInput: 0, expectedOutput: ""},
		{numInput: 1, expectedOutput: "1 "},
	}
	for _, test := range testCases {
		var buf bytes.Buffer
		out = &buf
		input := test.numInput
		fib(input)
		actual := ReplaceAllSpaceWithNewline(buf.String())
		expected := test.expectedOutput
		assert.Equalf(t, expected, actual, "Febonacci series failed.")
	}
}

func ReplaceAllSpaceWithNewline(temp string) string {
	return strings.ReplaceAll(temp, "\n", " ")
}
func TestMainOutput(t *testing.T) {
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`1
1
2
3
5
8
13
1
-1
2
-3
5
-8
13
`)
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual, "Unexpected output in main")
}
