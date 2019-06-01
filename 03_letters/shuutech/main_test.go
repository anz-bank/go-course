package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLetters(t *testing.T) {
	var m = letters("aba")
	var expected int
	var actual int
	for key, value := range m {
		if string(key) == "a" {
			expected = 2
			actual = value
			if expected != actual {
				t.Errorf("expected %v, got %v", actual, expected)
				t.Fail()
			}
		}
		if string(key) == "b" {
			expected = 1
			actual = value
			if expected != actual {
				t.Errorf("expected %v, got %v", actual, expected)
				t.Fail()
			}
		}
	}
}

func TestSortLetters(t *testing.T) {
	var m = letters("aba")
	var expected string
	var actual string
	var sorted = sortLetters(m)

	if sorted[0] != "a:2" {
		t.Errorf("expected %v, got %v", actual, expected)
		t.Fail()
	}
	if sorted[1] != "b:1" {
		t.Errorf("expected %v, got %v", actual, expected)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	main()
	fmt.Printf("test!")
	os.Exit(m.Run())
}
