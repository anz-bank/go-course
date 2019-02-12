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
	expected := "0\n1\n1\n2\n3\n5\n8\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestMinLimit(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	fib(-1)
	// Then
	expected := "Invalid Input\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestZeroSeries(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	fib(0)
	// Then
	expected := "Invalid Input\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestFirstSeries(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	fib(1)
	// Then
	expected := "0\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestSecondSeries(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	fib(2)
	// Then
	expected := "0\n1\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestHighValues(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	// When
	fib(20)
	// Then
	expected := "0\n1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n89\n144\n233\n377\n610\n987\n1597\n2584\n4181\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
