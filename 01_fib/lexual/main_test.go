package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{
			7,
			strings.Join([]string{"1", "1", "2", "3", "5", "8", "13"}, "\n") + "\n",
		},
		{
			-7,
			strings.Join([]string{"1", "-1", "2", "-3", "5", "-8", "13"}, "\n") + "\n",
		},
		{
			0,
			"",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		out = &buf

		fib(test.input)

		actual := buf.String()

		if test.expected != actual {
			// %q (raw-string), we don't want to print '\n' as newline.
			t.Errorf("\nActual: %q\nExpected: %q", actual, test.expected)
		}
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	expectedNums := []string{
		"1",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
	}

	main()

	actual := buf.String()
	expected := strings.Join(expectedNums, "\n") + "\n"

	if expected != actual {
		// %q (raw-string), we don't want to print '\n' as newline.
		t.Errorf("\nActual: %q\nExpected: %q", actual, expected)
	}
}
