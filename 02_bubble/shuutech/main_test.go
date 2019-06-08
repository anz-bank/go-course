package main

import (
	"fmt"
	"os"
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

func TestMain(m *testing.M) {
	main()
	fmt.Printf("test!")
	os.Exit(m.Run())
}
