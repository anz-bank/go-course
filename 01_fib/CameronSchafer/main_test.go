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

//FibOutput1 test
func TestFibOutput1(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(7)")
	}
}

//FibOutput2 test
func TestFibOutput2(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-7)

	expected := strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(7)")
	}
}

//NormalFibOutput test
func TestNormalFibOutput(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13}
	actual := calculateNormalFib(7)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected output in calculateNormalFib(7)")
	}
}

//NegaFibOutput test
func TestNegaFibOutput(t *testing.T) {
	calculateNegaFib(7)

	expected := []int{1, -1, 2, -3, 5, -8, 13}
	actual := calculateNegaFib(7)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected output in calculateNegaFib(7)")
	}
}

//CalcNextInSequence test
func TestCalcNextInSequence(t *testing.T) {
	expected := 5
	actual := calcNextInSequence(1, 4)

	if actual != expected {
		t.Errorf("Unexpected output in calcNextInSequence(1,4)")
	}
}

//PrintSequence test
func TestPrintSequence(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	printFibSequence([]int{1, 2, 3, 6, 23})

	expected := strconv.Quote("1\n2\n3\n6\n23\n")
	actual := strconv.Quote(buf.String())

	if actual != expected {
		t.Errorf("Unexpected output in printFibSequence([]int{1,2,3,6,23})")
	}
}
