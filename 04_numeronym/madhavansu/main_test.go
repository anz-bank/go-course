package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var s = []struct {
	input    string
	expected string
}{
	{"accessibility  ", "[a11y]"},
	{"Kubernetes", "[K8s]"},
	{"World Wide Web", "[W12b]"},
	{"மாதவ", "[ம2வ]"},
	{"ΩΟΞψωξ", "[Ω4ξ]"},
	{"  	   ", "[]"},
	{"This is a multi word checK", "[T24K]"},
	{"_This is an underscore and empty char          ", "[_35r]"},
	{"", "[]"},
	{"        +        ", "[+]"},
	{"        \n     \t   ", "[]"},
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
	// Given
	r := require.New(t)

	// When
	for _, v := range s {
		r.Equalf(v.expected, fmt.Sprint(numeronyms(v.input)), "Unexpected output in numeronyms()")
	}
}
