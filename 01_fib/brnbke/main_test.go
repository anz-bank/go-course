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

	expected := strconv.Quote(`0
1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected %v, sadly got %v", expected, actual)
	}
}

func TestFib(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  string
	}{
		"-4": {input: -4, want: `0
1
-1
2
-3
`},
		"3": {input: 3, want: `0
1
1
2
`},
		"0": {input: 0, want: `0
`},
		"93": {input: 93, want: `Fibonacci numbers greater than 92 not supported
`},
		"-93": {input: -93, want: `Fibonacci numbers greater than 92 not supported
`},
	}

	for name, test := range testCases {
		test := test
		var buf bytes.Buffer
		out = &buf
		t.Run(name, func(t *testing.T) {
			fib(test.input)
			result := buf.String()
			if result != test.want {
				t.Errorf("expected %v, got %v", test.want, result)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  int
	}{
		"positive": {input: 2, want: 2},
		"negative": {input: -3, want: 3},
		"zero":     {input: 0, want: 0},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := abs(test.input)
			if result != test.want {
				t.Errorf("expected %v, got %v", test.want, result)
			}
		})
	}
}
