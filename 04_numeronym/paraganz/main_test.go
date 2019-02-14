package main

import (
	"bytes"
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
	{"abcd", []string{"a2d"}},
	{"a", []string{"a"}},
	{"", []string{""}},
	{" ", []string{" "}},
	{"abc⌘ef", []string{"a4f"}},
	{" abcd", []string{"a2d"}},
	{"⌘abcef❤", []string{"⌘5❤"}},
	{"abcd ", []string{"a2d"}},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronymsOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		out := numeronyms(tt.in)
		r.ElementsMatch(tt.out, out)
	}
}
