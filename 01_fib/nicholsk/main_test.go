package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(4)

	expected := strconv.Quote("1\n1\n2\n3\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(4)")
	}

}

func TestFibWithZero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("0\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(0)")
	}
}

func TestFibWithNegative(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-1)

	expected := strconv.Quote("fib(n) doesn't accept negative integers\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(0)")
	}
}
