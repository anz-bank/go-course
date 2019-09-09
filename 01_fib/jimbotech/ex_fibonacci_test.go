package main

import (
	"bytes"
	"testing"
)

func TestFib(t *testing.T) {

	fibResults := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

	f := fibonacci()

	for _, val := range fibResults {
		if r := f(); r != val {
			t.Errorf("returned value %v does not match %v", r, val)
		}
	}
}

func TestMain(t *testing.T) {

	want := "1\n1\n2\n3\n5\n8\n13\n"
	var buf bytes.Buffer
	out = &buf
	main()
	result := buf.String()

	if result != want {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestNeg(t *testing.T) {

	want := "1\n-1\n2\n-3\n5\n-8\n13\n"
	var buf bytes.Buffer
	out = &buf
	fib(-7)
	result := buf.String()

	if result != want {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestZero(t *testing.T) {

	var buf bytes.Buffer
	out = &buf
	fib(0)
	result := buf.String()

	if len(result) > 0 {
		t.Errorf("expected nothing to be printed, got %v", result)
	}
}

func TestOne(t *testing.T) {

	want := "1\n"
	var buf bytes.Buffer
	out = &buf
	fib(1)
	result := buf.String()

	if result != want {
		t.Errorf("expected %v, got %v", want, result)
	}
}
