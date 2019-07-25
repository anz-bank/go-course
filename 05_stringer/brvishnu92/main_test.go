package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := `127.0.0.1
`
	actual := buf.String()
	assert.Equalf(t, expected, actual, "Got: %v Want: %v", actual, expected)
}

func TestStringer(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"normal":   {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		"empty":    {input: IPAddr{}, expected: "0.0.0.0"},
		"oneval":   {input: IPAddr{127}, expected: "127.0.0.0"},
		"threeval": {input: IPAddr{127, 12, 122}, expected: "127.12.122.0"},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprint(test.input)
			assert.Equalf(t, test.expected, result, "Got: %v Want: %v", result, test.expected)
		})
	}
}
