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

	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}

func TestFibNums(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	tests := []struct {
		input int
		want  string
	}{
		{input: -1, want: strconv.Quote("")},
		{input: 0, want: strconv.Quote("")},
		{input: 1, want: strconv.Quote("1\n")},
		{input: 7, want: strconv.Quote("1\n1\n1\n2\n3\n5\n8\n13\n")},
	}

	for _, test := range tests {
		fib(test.input)
		got := strconv.Quote(buf.String())
		if test.want != got {
			t.Fatalf("Expected: %v, Got: %v", test.want, got)
		}
	}
}
