package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestBubble(t *testing.T) {
	var unsorted = []int{3, 2, 1, 5}
	var sorted = bubble(unsorted)
	var s = ""

	for i := range sorted {
		s = s + "," + strconv.Itoa(sorted[i])
	}
	var expected = ",1,2,3,5"
	if expected != s {
		t.Errorf("expected %v, got %v", expected, s)
		t.Fail()
	}
}

func TestBubbleOneLength(t *testing.T) {
	var unsorted = []int{1}
	var sorted = bubble(unsorted)
	var s = ""

	for i := range sorted {
		s = s + "," + strconv.Itoa(sorted[i])
	}
	var expected = ",1"
	if expected != s {
		t.Errorf("expected %v, got %v", expected, s)
		t.Fail()
	}
}

func TestBubbleZeroLength(t *testing.T) {
	var unsorted = []int{}
	var sorted = bubble(unsorted)
	var s = ""

	for i := range sorted {
		s = s + "," + strconv.Itoa(sorted[i])
	}
	var expected = ""
	if expected != s {
		t.Errorf("expected %v, got %v", expected, s)
		t.Fail()
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[1 2 3 5]"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected: %q, actual: %q", expected, actual)
	}
}
