package main

import (
	"bytes"
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
	want := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, want, actual, "Numeronym failed.")
}

var testCases = map[string]struct {
	input,
	want []string
}{
	"main":               {input: []string{"accessibility", "Kubernetes", "abc"}, want: []string{"a11y", "K8s", "abc"}},
	"Small strings":      {input: []string{"x", "yz", "xyz", "wxyz"}, want: []string{"x", "yz", "xyz", "w2z"}},
	"String lines":       {input: []string{"Have fun", "Go Lang"}, want: []string{"H6n", "G5g"}},
	"Alphanumeric":       {input: []string{"12as34sdf5678asd90", "abcd232ef"}, want: []string{"1160", "a7f"}},
	"Empty":              {input: []string{"", ""}, want: []string{"", ""}},
	"Special Characters": {input: []string{"%*(*)ASDjahs(", "absd*?>^&%$asd"}, want: []string{"%11(", "a12d"}},
	"Unicode":            {input: []string{"aübenäc", "übeäcgä", "üää"}, want: []string{"a5c", "ü5ä", "üää"}},
}

func TestNumeronym(t *testing.T) {
	for caseName, test := range testCases {
		want := test.want
		numeronym := numeronyms(test.input...)
		assert.Equalf(t, want, numeronym, "Numeronyms function failed for TC: %v", caseName)
	}
}
