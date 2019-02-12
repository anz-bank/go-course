package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibinocci(t *testing.T) {
	r := require.New(t)

	r.Equal(1, fibinocci(2))
	r.Equal(2, fibinocci(3))
	r.Equal(0, fibinocci(0))
	r.Equal(-1, fibinocci(-1))
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "1\n1\n2\n3\n5\n8\n13\n21\n34\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
