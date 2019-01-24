package main

import (
	"bytes"
	"strings"
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
	expected := "Hallo du schöne Welt!\\n" // escape newline for easier string comparison
	actual := strings.Replace(buf.String(), "\n", "\\n", -1)
	r.Equalf(expected, actual, "Unexpected output in main()")
}
