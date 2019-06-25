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

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")

	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output from main(), expected = %v, actual = %v", expected, actual)
	}
}

func TestFibNums(t *testing.T) {

	tests := map[string]struct {
		input int
		want  string
	}{
		"fib(-1)":  {input: -1, want: strconv.Quote("1\n")},
		"fib(0)":   {input: 0, want: strconv.Quote("0\n")},
		"fib(1)":   {input: 1, want: strconv.Quote("1\n")},
		"fib(-7)":  {input: -7, want: strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n")},
		"fib(7)":   {input: 7, want: strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")},
		"fib(93)":  {input: 93, want: strconv.Quote("Range for fib(n) should be between -92 to 92\n")},
		"fib(-93)": {input: -93, want: strconv.Quote("Range for fib(n) should be between -92 to 92\n")},
	}

	for name, test := range tests {
		var buf bytes.Buffer
		out = &buf
		test := test

		t.Run(name, func(t *testing.T) {
			fib(test.input)
			got := strconv.Quote(buf.String())
			if test.want != got {
				t.Fatalf("Test Expected: %v, Got: %v", test.want, got)
			}
		})
	}
}
