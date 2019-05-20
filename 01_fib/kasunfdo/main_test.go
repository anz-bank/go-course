package main

import (
	"bytes"
	"fmt"
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

func TestFibZero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("0\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibPosOne(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote("1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibNegOne(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-1)

	expected := strconv.Quote("1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
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

	fmt.Println(expected)
	fmt.Println(actual)
	fmt.Println(expected != actual)
	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
