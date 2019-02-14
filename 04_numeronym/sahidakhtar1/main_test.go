package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutPut(t *testing.T) {
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

func TestNumeronyms(t *testing.T) {
	testData := []struct {
		in  []string
		out []string
	}{
		{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
		{[]string{"accessibility", "Kubernetes", "abc", "accessibility"}, []string{"a11y", "K8s", "abc", "a11y"}},
		{[]string{"123456"}, []string{"146"}},
		{[]string{"-123456"}, []string{"-56"}},
		{[]string{"aaaaaaaaaaaaaaaa"}, []string{"a14a"}},
		{[]string{"  accessibility  ", "Kubernetes  ", "  abc"}, []string{"a11y", "K8s", "abc"}},
	}
	r := require.New(t)
	for _, tt := range testData {
		result := numeronyms(tt.in...)
		r.Equalf(tt.out, result, "Numeronymes are not right")
	}
}

func TestCreateNumeronym(t *testing.T) {
	testData := []struct {
		in  string
		out string
	}{
		{"accessibility", "a11y"},
		{"Kubernetes", "K8s"},
		{"abc", "abc"},
		{"ab", "ab"},
		{"b", "b"},
		{"", ""},
		{"  accessibility  ", "a11y"},
		{"  Kubernetes", "K8s"},
		{"abc  ", "abc"},
		{"abc efhjb", "abc efhjb"},
		{"abc\\tefhjb", "a8b"},
		{"abc\nefhjb", "abc\nefhjb"},
		{"abcefhjb\n", "a6b"},
		{"世🖖界", "世🖖界"},
		{"世", "世"},
		{"a世🖖界bc", "a4c"},
		{"世abc", "世2c"},
		{"abc世", "a2世"},
		{"⌘abcef♥", "⌘5♥"},
	}
	r := require.New(t)
	for _, tt := range testData {
		result := createNumeronym(tt.in)
		r.Equalf(tt.out, result, "Numeronymes are not right")
	}
}
