package main

import (
	"bytes"
	"strconv"
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
	expected := strconv.Quote("Hallo du schöne Welt!\n1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
