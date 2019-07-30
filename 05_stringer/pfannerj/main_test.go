package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf
	main()
	expected := "127.0.0.1\n"
	actual := buf.String()
	if expected != actual {
		t.Errorf(expected, actual, "Unexpected output in main()")
	}
}

var IPAddrCases = map[string]struct {
	input    IPAddr
	expected string
}{
	"StdLocalHost": {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
	"Empty":        {input: IPAddr{}, expected: "0.0.0.0"},
	"Zeroes":       {input: IPAddr{0, 0, 0, 0}, expected: "0.0.0.0"},
	"One":          {input: IPAddr{15}, expected: "15.0.0.0"},
	"Two":          {input: IPAddr{152, 73}, expected: "152.73.0.0"},
	"Three":        {input: IPAddr{5, 89, 176}, expected: "5.89.176.0"},
	"Random":       {input: IPAddr{1, 12, 234, 9}, expected: "1.12.234.9"},
	"Max":          {input: IPAddr{255, 255, 255, 255}, expected: "255.255.255.255"},
}

func TestIPAddr(t *testing.T) {
	for name, tc := range IPAddrCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := tc.input.String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
