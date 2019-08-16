package main

import (
	"bytes"

	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	buf := bytes.Buffer{}
	out = &buf
	main()
	expected := strconv.Quote("127.0.0.1")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestStringer(t *testing.T) {
	testCases := []struct {
		input    IPAddr
		expected string
	}{
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{IPAddr{}, "0.0.0.0"},
		{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
		{IPAddr{10, 74, 83, 7}, "10.74.83.7"},
	}
	for _, testCase := range testCases {
		if !assert.Equal(t, testCase.input.String(), testCase.expected) {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", testCase.input, testCase.expected)
		}
	}
}
