package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestNumeronymsOutput(t *testing.T) {

	expected := []string{}
	actual := numeronyms()

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}

	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("Unexpected output in main()")
		}
	}
}

func TestNumeronymsOutput2(t *testing.T) {

	expected := []string{"a", "ab", "abc", "a2d", "a3e"}
	actual := numeronyms("a", "ab", "abc", "abcd", "abcde")

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}

	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("Unexpected output in main()")
		}
	}
}
