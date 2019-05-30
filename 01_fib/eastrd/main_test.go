package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestFibOutput1(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutput2(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote("1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutput3(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("Fibonacci number has to be positive\n")
	actual := strconv.Quote(buf.String())

	fmt.Println(actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutput4(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(4)

	expected := strconv.Quote("1\n1\n2\n3\n")
	actual := strconv.Quote(buf.String())

	fmt.Println(actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutput5(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(2)

	expected := strconv.Quote("1\n1\n")
	actual := strconv.Quote(buf.String())

	fmt.Println(actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutput6(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("Fibonacci number has to be positive\n")
	actual := strconv.Quote(buf.String())

	fmt.Println(actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n")
	actual := strconv.Quote(buf.String())

	fmt.Println(actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
