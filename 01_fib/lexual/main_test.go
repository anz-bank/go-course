package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFibPositive(t *testing.T) {
	numbers := []string{
		"1",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
	}
	runTest(t, 7, numbers)
}

func TestFibNegative(t *testing.T) {
	numbers := []string{
		"1",
		"-1",
		"2",
		"-3",
		"5",
		"-8",
		"13",
	}
	runTest(t, -7, numbers)
}

func runTest(t *testing.T, n int, expectedNums []string) {
	var buf bytes.Buffer
	out = &buf

	fib(n)

	actual := buf.String()
	expected := strings.Join(expectedNums, "\n") + "\n"

	if expected != actual {
		// %q (raw-string), we don't want to print '\n' as newline.
		t.Errorf("\nActual: %q\nExpected: %q", actual, expected)
	}
}
