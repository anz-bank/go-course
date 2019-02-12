package main

import (
	"bytes"
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
	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestFibonacciNegativeOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fibonacciSeries(-5)

	// Then
	expected := "1\n-1\n2\n-3\n5\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestFibonacciZeroOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fibonacciSeries(0)

	// Then
	expected := ""
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
