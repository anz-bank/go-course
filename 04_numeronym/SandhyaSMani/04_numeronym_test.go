package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in  string
	out []string
}{
	{"numeronym", []string{"n7m"}},
	{"numeronym test", []string{"n12t"}},
	{"abc", []string{"abc"}},
	{"abcd", []string{"a2d"}},
	{"ab", []string{"ab"}},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronymsOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		out := numeronyms(tt.in)
		r.ElementsMatch(tt.out, out)
	}
}
