package main

import (
	"testing"
)

func TestBbl1(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := bubbleSort([]int{5, 4, 3, 2, 1})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubbleSort function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubbleSort function")
		}
	}
}

func TestBbl2(t *testing.T) {
	expected := []int{}
	actual := bubbleSort([]int{})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubbleSort function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubbleSort function")
		}
	}
}

func TestBbl3(t *testing.T) {
	expected := []int{999999}
	actual := bubbleSort([]int{999999})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubbleSort function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubbleSort function")
		}
	}
}

func TestBbl4(t *testing.T) {
	expected := []int{0, 0, 1, 1}
	actual := bubbleSort([]int{1, 0, 1, 0})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubbleSort function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubbleSort function")
		}
	}
}

func TestBbl5(t *testing.T) {
	expected := []int{-9, -5, -4, 0}
	actual := bubbleSort([]int{-5, 0, -4, -9})

	if (expected == nil) != (actual == nil) || len(expected) != len(actual) {
		t.Errorf("Error in bubbleSort function")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Error in bubbleSort function")
		}
	}
}
