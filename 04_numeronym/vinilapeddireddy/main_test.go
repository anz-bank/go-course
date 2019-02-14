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
	{[]string{}, []string{}},
	{[]string{"accessibility", "Kubernetes", "abc", "charecters"}, []string{"a11y", "K8s", "abc", "c8s"}},
	{[]string{"!!&^#*&+"}, []string{"!6+"}},
	{[]string{"🦋🦋!&^#*&+ 🦋"}, []string{"🦋9🦋"}},
	{[]string{"a1y", "K8s", "a1c", "c8s"}, []string{"a1y", "K8s", "a1c", "c8s"}},
}

func TestNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		input := tt.inputStrings
		output := numeronyms(input...)
		r.Equal(tt.outputStrings, output)
	}
}

func TestNumeronymsMainOutput(t *testing.T) {
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
