package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

//main test
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

//FibOutput test
func TestFibOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(7)")
	}
}

//NormalFibOutput test
func TestNormalFibOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	expected := []int{1, 1, 3, 5, 8, 13}
	actual := calculateNormalFib(7)

	if reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected output in calculateNormalFib(7)")
	}
}

//NegaFibOutput test
func TestNegaFibOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	calculateNegaFib(7)

	expected := []int{1, -1, 3, -5, 8, -13}
	actual := calculateNegaFib(7)

	if reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected output in calculateNegaFib(7)")
	}
}
