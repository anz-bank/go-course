package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData = []struct {
	in       []string
	expected []string
}{
	{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
	{[]string{"a", "bbbb", "abc", "abcd"}, []string{"a", "b2b", "abc", "a2d"}},
	{[]string{"cabc ", "b ab", "abc", "abcd"}, []string{"c2c", "b2b", "abc", "a2d"}},
	{[]string{"你和我", "你和我s"}, []string{"你和我", "你2s"}},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	r.Equalf(`[a11y K8s abc]`, buf.String(), "Unexpected output in main()")
}

func TestNumeronyms(t *testing.T) {
	for _, v := range testData {
		//When
		l := v.in

		//Then
		require.Equal(t, v.expected, numeronyms(l...), "Unexpected output in numeronyms()")
	}
}
