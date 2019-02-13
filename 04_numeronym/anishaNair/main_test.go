package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	inputStrings  []string
	outputStrings []string
}{
	{[]string{}, nil},
	{[]string{"accessibility", "Kubernetes", "abc"},
		[]string{"a11y", "K8s", "abc"}},
	{[]string{"!hdfdfgjdhf djfhdfh"},
		[]string{"!17h"}},
}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	main()
	// Then
	expected := `[a11y K8s abc]
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		input := []string(tt.inputStrings)
		output := getNumeronyms(input...)
		r.Equal(tt.outputStrings, output)
	}
}
