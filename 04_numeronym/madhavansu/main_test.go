package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var s = []struct {
	input  string
	output string
}{
	{"accessibility", "[a11y]"},
	{"Kubernetes", "[K8s]"},
	{"World Wide Web", "[W12b]"},
	{"abc", "[abc]"},
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`[a11y K8s abc]`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestNumeronyms(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	for _, v := range s {
		expected := fmt.Sprint(numeronyms(v.input))
		actual := v.output
		r.Equalf(expected, actual, "Unexpected output in numeronyms()")
	}
}
