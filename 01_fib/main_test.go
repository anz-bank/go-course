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
	outWriter = &buf

	// When
	main()

	// Then
	expected := `1
1
2
3
5
8
13`
	actual := buf.String()
	r.Equalf(expected, strings.TrimSpace(actual), "Unexpected output in main()")
}

func TestNegativeNumber(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	fib(-1)

	actual := buf.String()
	r.Equalf("", actual, "Unexpected output for negative number test")
}

func TestZero(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	fib(0)

	actual := buf.String()
	r.Equalf("", actual, "Unexpected output for test with 0")
}
func TestForOne(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	outWriter = &buf

	// When
	fib(1)

	// Then
	expected := `1`
	actual := buf.String()
	r.Equalf(expected, strings.TrimSpace(actual), "Unexpected output for test with one")
}
