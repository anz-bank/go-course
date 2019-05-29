package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf

	main()

	expected := strconv.Quote("Fibonacci series completed...\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutputFibZero(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(0)

	expected := strconv.Quote("")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(0)")
	}
}

func TestFibOutputFibSeven(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(7)")
	}
}

func TestFibOutputFibNine(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(9)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n21\n34\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(9)")
	}
}

func TestFibOutputNegaFibNine(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(-9)

	expected := strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n-21\n34\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(-9)")
	}
}
