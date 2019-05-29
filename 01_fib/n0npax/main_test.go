package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	main()

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

func TestFibOutput(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	fib(6)

	expected := strconv.Quote(`1
1
2
3
5
8
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in fib()")
}

func TestFibOutputNegative(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	fib(-7)
	expected := strconv.Quote(`1
-1
2
-3
5
-8
13
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in fib()")
}

func TestFibOutputWithZero(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	fib(0)
	expected := strconv.Quote(`0
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in fib()")
}

func TestFibOutputWithOne(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	fib(1)
	expected := strconv.Quote(`1
`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in fib()")
}
