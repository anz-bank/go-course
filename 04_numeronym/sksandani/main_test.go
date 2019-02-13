package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	inputString      []string
	numeronymsOutput []string
}{
	{[]string{}, []string{}},
	{[]string{"", "+12345", "-1234"}, []string{"", "+45", "-34"}},
	{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}}}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := string("[a11y K8s abc]\n")
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		outMap := numeronyms(tt.inputString...)
		r.Equal(tt.numeronymsOutput, outMap)
	}
}
