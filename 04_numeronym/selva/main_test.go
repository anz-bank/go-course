package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var nemeronymsTest = []struct {
	in  []string
	out []string
}{
	{[]string{"accessibility", "internationalisation"}, []string{"a11y", "i18n"}},
}

func TestMainTableDriven(t *testing.T) {
	r := require.New(t)
	for _, tt := range nemeronymsTest {
		out := nemeronyms(tt.in...)
		r.EqualValues(tt.out, out, "Unexpected output in main()")
	}
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[a11y K8s abc a2d]")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
