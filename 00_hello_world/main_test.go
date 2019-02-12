package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//When
	main()

	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Fib () doesn't work for 7")

}

func TestFibFor7(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//When
	fib(7)

	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Fib () doesn't work for 7")

}

func TestFibForZero(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//When
	fib(0)

	expected := ""
	actual := buf.String()
	r.Equalf(expected, actual, "Fib () doesn't for Zero")

}
