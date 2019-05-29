package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFib1(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote(`1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Helper()
		t.Errorf("Unexpected output in main()")
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
		t.Errorf(actual)
	}
}
