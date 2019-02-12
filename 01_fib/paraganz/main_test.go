package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestFibOutputWithZero(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fib(0)

	//Then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}

func TestFibOutputWithNegativeOne(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fib(-1)

	//Then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}

func TestFibOutputWithOne(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fib(1)

	//Then
	r.Equalf("1\n", buf.String(), "Unexpected output in main()")
}
