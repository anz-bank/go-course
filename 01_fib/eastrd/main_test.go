package main

import (
	"testing"
)

func TestFib1(t *testing.T) {

	expected := 5
	actual, err := fib(5)

	if expected != actual && err != nil {
		t.Errorf("Error in fib function")
	}
}

func TestFib2(t *testing.T) {

	expected := 1
	actual, err := fib(1)

	if expected != actual && err != nil {
		t.Errorf("Error in fib function")
	}
}

func TestFib3(t *testing.T) {

	expected := 1
	actual, err := fib(2)

	if expected != actual && err != nil {
		t.Errorf("Error in fib function")
	}
}

func TestFib4(t *testing.T) {

	_, err := fib(-1)

	if err == nil {
		t.Errorf("Error in fib function")
	}
}
