package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("127.0.0.1")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "IP Address formatting failed.")
}

var testCases = map[string]struct {
	input IPAddr
	want  string
}{
	"main":            {input: IPAddr{127, 0, 0, 1}, want: "127.0.0.1"},
	"empty":           {input: IPAddr{}, want: "0.0.0.0"},
	"single":          {input: IPAddr{1}, want: "1.0.0.0"},
	"partial":         {input: IPAddr{192, 168}, want: "192.168.0.0"},
	"All IPs":         {input: IPAddr{0, 0, 0, 0}, want: "0.0.0.0"},
	"Special Address": {input: IPAddr{255, 255, 255, 255}, want: "255.255.255.255"},
}

func TestIPAddr(t *testing.T) {
	for caseName, test := range testCases {
		input, expected := test.input, test.want
		actual := fmt.Sprint(input)
		assert.Equalf(t, expected, actual, "IP Address formatting failed for TC: %v", caseName)
	}
}
