package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFibonacciForPositiveNumber(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(3)
	expected := strconv.Quote("1\n1\n2\n")
	actual := strconv.Quote(buf.String())
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestFibonacciForNegativeNumber(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(-3)
	expected := strconv.Quote("-1\n-1\n-2\n")
	actual := strconv.Quote(buf.String())
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestFibonacciForZeroNumber(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(0)
	expected := strconv.Quote("")
	actual := strconv.Quote(buf.String())
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
