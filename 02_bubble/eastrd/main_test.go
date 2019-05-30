package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestBbl1(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := bubble([]int{5, 4, 3, 2, 1})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubble function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubble function")
		}
	}
}

func TestBbl2(t *testing.T) {
	expected := []int{}
	actual := bubble([]int{})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubble function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubble function")
		}
	}
}

func TestBbl3(t *testing.T) {
	expected := []int{999999}
	actual := bubble([]int{999999})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubble function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubble function")
		}
	}
}

func TestBbl4(t *testing.T) {
	expected := []int{0, 0, 1, 1}
	actual := bubble([]int{1, 0, 1, 0})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubble function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubble function")
		}
	}
}

func TestBbl5(t *testing.T) {
	expected := []int{-9, -5, -4, 0}
	actual := bubble([]int{-5, 0, -4, -9})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubble function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubble function")
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Error in bubble function")
	}
}
