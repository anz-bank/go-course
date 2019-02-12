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
	expected := `[1 2 3 5]`
	actual := strings.TrimSuffix(buf.String(), "\n")
	r.Equalf(expected, actual, "Unexpected output in main()")
}