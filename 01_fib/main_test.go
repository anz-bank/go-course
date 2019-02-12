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
	expected := `1
1
2
3
5
8
13
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestFibOutputForNegafibonacciNumbers(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fib(-7)

	// Then
	expected := `13
-8
5
-3
2
-1
1
`
	actual := buf.String()
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
func TestFibOutputWithMinusOne(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fib(-1)

	//Then
	r.Equalf("1\n", buf.String(), "Unexpected output in main()")
}
