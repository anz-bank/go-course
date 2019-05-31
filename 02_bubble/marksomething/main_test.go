package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	actual := buf.String()
	expected := "[1 2 3 5]\n"
	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}

func TestBubble(t *testing.T) {
	actual := fmt.Sprint(bubble([]int{3, 2, 1, 5, 3, 2, 1}))
	expected := "[1 1 2 2 3 3 5]"
	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}

func TestBubbleNegative(t *testing.T) {
	actual := fmt.Sprint(bubble([]int{3, -2, 1, 5, -3, 2, 1}))
	expected := "[-3 -2 1 1 2 3 5]"
	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}

func TestBubbleNoMutation(t *testing.T) {
	a := []int{3, 1, 2}
	bubble(a)
	expected := "[3 1 2]"
	actual := fmt.Sprint(a)

	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}

func TestBubbleSingleValue(t *testing.T) {
	actual := fmt.Sprint(bubble([]int{1}))
	expected := "[1]"
	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}

func TestBubbleEmpty(t *testing.T) {
	actual := fmt.Sprint(bubble([]int{}))
	expected := "[]"
	if actual != expected {
		t.Errorf("Actual (%q) does not equal Expected (%q)", actual, expected)
	}
}
