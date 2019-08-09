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
	expected := "127.0.0.1"
	actual := buf.String()
	assert.Equal(t, expected, actual, "Unexpected output in main() actual %v but expected %v", actual, expected)
}
func TestIPAddrStringer(t *testing.T) {

	testCases := []struct {
		description string
		input       IPAddr
		expected    string
	}{
		{description: "Empty", input: IPAddr{}, expected: "0.0.0.0"},
		{description: "All zeroes", input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
		{description: "Single octet", input: IPAddr{115}, expected: "115.0.0.0"},
		{description: "Two octets", input: IPAddr{115, 70}, expected: "115.70.0.0"},
		{description: "Three octets", input: IPAddr{115, 70, 166}, expected: "115.70.166.0"},
		{description: "IP Address", input: IPAddr{115, 70, 166, 151}, expected: "115.70.166.151"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.description, func(t *testing.T) {
			actual := test.input.String()
			expected := test.expected
			assert.Equal(t, actual, expected, "actual %v but expected %v", actual, expected)
		})
	}

}
