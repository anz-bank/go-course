package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var tests = []struct {
	in  []string
	out []string
}{
	{[]string{"internationalization", "Administration", "Africanization", "Appocalypse"}, []string{"i18n", "A12n",
		"A12n", "A9e"}},
	{[]string{}, []string{}},
}

func TestNumeronymsWithValues(t *testing.T) {
	r := require.New(t)

	for _, tt := range tests {
		out := numeronyms(tt.in...)
		r.EqualValues(tt.out, out)
	}
}
