package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testMatrix = []struct {
	input  []string
	output []string
}{
	{[]string{"abba", "accessibility", "Kubernetes", "abc"}, []string{"a2a", "a11y", "K8s", "abc"}},
	{[]string{"abb a"}, []string{"a3a"}},
	{[]string{""}, []string{""}},
	{[]string{"£€€€§‡®"}, []string{"£5®"}},
	{[]string{"ßäöüÄÖÜ"}, []string{"ß5Ü"}}}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `[a11y K8s abc]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronyms(t *testing.T) {
	// Given
	r := require.New(t)
	for _, testData := range testMatrix {
		numeronymsVal := numeronyms(testData.input...)
		r.Equal(testData.output, numeronymsVal)
	}
}
