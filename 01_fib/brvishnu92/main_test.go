package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)
func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")
	}

}

func TestMainOutputPositive(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")
	}

}

func TestMainOutputNegative(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-7)

	expectedNeg := strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n")
	actualNeg := strconv.Quote(buf.String())

	if expectedNeg != actualNeg {
		fmt.Println(actualNeg)
		t.Errorf("Unexpected output in main()")
	}

}

func TestMainOutputZero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expectedNeg := strconv.Quote("0\n")
	actualNeg := strconv.Quote(buf.String())

	if expectedNeg != actualNeg {
		fmt.Println(actualNeg)
		t.Errorf("Unexpected output in main()")
	}

}
