package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFibNeg11(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-11)

	expected := strconv.Quote(``)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Helper()
		t.Errorf(expected)
		t.Errorf(actual)
	}
}
func TestFib0(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote(``)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Helper()
		t.Errorf(expected)
		t.Errorf(actual)
	}
}
func TestFib1(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote(`1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Helper()
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestFib7(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`Call Fibonacci numbers!
1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Helper()
		t.Errorf(expected)
		t.Errorf(actual)
	}
}
