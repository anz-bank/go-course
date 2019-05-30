package main

import (
	"bytes"
	"testing"
)

func TestFibPos(t *testing.T) {
	var buf_fib bytes.Buffer
	out_fib = &buf_fib

	fib(7)

	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buf_fib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
func TestFibNeg(t *testing.T) {
	var buf_fib bytes.Buffer
	out_fib = &buf_fib

	fib(-7)

	expected := "1\n-1\n2\n-3\n5\n-8\n13\n"
	actual := buf_fib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
