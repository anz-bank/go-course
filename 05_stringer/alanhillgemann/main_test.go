package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("127.0.0.1")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
}

func TestStringer(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"Nil":        {input: IPAddr{}, expected: "0.0.0.0"},
		"Incomplete": {input: IPAddr{1}, expected: "1.0.0.0"},
		"All 0":      {input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
		"Variable":   {input: IPAddr{4, 3, 2, 1}, expected: "4.3.2.1"},
		"Max":        {input: IPAddr{255, 255, 255, 255}, expected: "255.255.255.255"},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := fmt.Sprint(test.input)
			expected := test.expected
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}
