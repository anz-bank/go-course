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
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestFibZero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("0\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestFibPosOne(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote("1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestFibNegOne(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-1)

	expected := strconv.Quote("1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestFibNegative(t *testing.T) {
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

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
