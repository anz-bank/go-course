package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
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

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestZeroInput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("Invalid Input: Must be a positive integer\n")

	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output: error message expected")
	}
}

func TestNegativeInput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-3)

	expected := strconv.Quote("Invalid Input: Must be a positive integer\n")

	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output: error message expected")
	}
}

func TestOneInput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote(`1
`)

	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
