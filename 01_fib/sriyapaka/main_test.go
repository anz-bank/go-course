package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFib(t *testing.T) {

	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("unexpected result - fib function")
	}
}

func TestNegativeSeries(t *testing.T) {

	var buf bytes.Buffer
	out = &buf
	fib(-7)
	expected := strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("unexpected result for negative series")
	}
}
