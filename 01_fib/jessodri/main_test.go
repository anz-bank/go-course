package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(format("1 1 2 3 5 8 13 "))
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("actual: %s does not match expected: %s", actual, expected)
	}
}

func TestFib(t *testing.T) {
	testCases := []struct {
		num    int
		result string
	}{
		{0, ""},
		{1, "1 "},
		{7, "1 1 2 3 5 8 13 "},
		{-1, "1 "},
		{-7, "1 -1 2 -3 5 -8 13 "},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer
		out = &buf
		fib(tc.num)
		expected := strconv.Quote(tc.result)
		actual := strconv.Quote(strings.Replace(buf.String(), "\n", " ", -1))
		if expected != actual {
			t.Errorf("Fibonacci sequence (%d) was incorrect, actual: %s, expected: %s.", tc.num, actual, expected)
		}
	}
}

func format(fib string) string {
	return strings.Replace(fib, " ", "\n", -1)
}
