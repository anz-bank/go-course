package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib

	main()

	expected := "1\n1\n2\n3\n5\n8\n13\n1\n-1\n2\n-3\n5\n-8\n13\n"
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
func TestFibPos(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib

	fib(7)

	expected := "1\n1\n2\n3\n5\n8\n13\n"
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
func TestFibNeg(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib

	fib(-7)

	expected := "1\n-1\n2\n-3\n5\n-8\n13\n"
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
