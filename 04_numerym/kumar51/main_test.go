package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumeroniumOutput(t *testing.T) {
	// Given
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

