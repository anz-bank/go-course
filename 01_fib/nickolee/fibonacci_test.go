package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestFibonacci(t *testing.T) {
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
		t.Errorf("TestFibonacci: Unexpected output in main()")
	}
}

func TestFibonacciUpperLimit(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(93)
	expected := strconv.Quote("Overflow error. Please use an argument less than 92")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("TestFibonacciUpperLimit: Unexpected output in main()")
	}

}

func TestFibonacciZero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(0)
	expected := strconv.Quote("0")
	actual := strconv.Quote(buf.String())

	// not allowed to compare equality for two slices. Hence used reflect.DeepEqual()
	// referred to https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFibonacciZero: Unexpected output in main()")
	}

}
