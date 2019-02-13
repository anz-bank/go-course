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
	{"accessibility", []string{"a11y"}},
	{"accessibility testing", []string{"a19g"}},
	{"abc", []string{"abc"}},
	{"a", []string{"a"}},
	{"", []string{""}},
	{" ", []string{" "}},
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
