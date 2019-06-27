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

	expected := strconv.Quote("0\n1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected result in main()")
	}
}

func TestFibNegativeOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	if fib(-7) != nil {
		t.Errorf("ufib(-7) unexpectedly threw an error")
	}

	expected := strconv.Quote("0\n1\n-1\n2\n-3\n5\n-8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected result in negative fib")
	}
}

func TestOverflowError(t *testing.T) {
	if fib(61) == nil {
		t.Errorf("fib(61) should have returned an error")
	}
}

func TestUnderflowError(t *testing.T) {
	if fib(-61) == nil {
		t.Errorf("fib(-61) should have returned an error")
	}
}

var fibloopTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFibonacciLoop(t *testing.T) {
	for _, tt := range fibloopTests {
		actual := fibonacciLoop(tt.n)
		if actual != tt.expected {
			t.Errorf("fibonacciLoop(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{-1, 1},
	{-2, -1},
	{-3, 2},
	{-4, -3},
	{-5, 5},
	{-6, -8},
	{-7, 13},
}

func TestFibonacci(t *testing.T) {
	for _, tt := range fibTests {
		actual := fibonacci(tt.n)
		if actual != tt.expected {
			t.Errorf("fibonacci(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

var signTests = []struct {
	n        int // input
	expected int // expected result
}{
	{-6, -1},
	{-4, -1},
	{-3, 1},
	{6, 1},
	{2, 1},
}

func TestFibonacciSign(t *testing.T) {
	for _, tt := range signTests {
		actual := fibonacciSign(tt.n)
		if actual != tt.expected {
			t.Errorf("fibonacciSign(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

var absTests = []struct {
	n        int // input
	expected int // expected result
}{
	{-6, 6},
	{-4, 4},
	{-3, 3},
	{6, 6},
	{2, 2},
}

func TestAbs(t *testing.T) {
	for _, tt := range absTests {
		actual := abs(tt.n)
		if actual != tt.expected {
			t.Errorf("abs(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
