package main

import (
	"bytes"
	"strconv"
	"testing"
)

func Test_abs(t *testing.T) {
	testCases := map[string]struct {
		input  int
		output int
	}{
		"normal test 1": {
			input:  1,
			output: 1,
		},
		"normal test 2": {
			input:  -1,
			output: 1,
		},
		"normal test 3": {
			input:  0,
			output: 0,
		},
	}

	for testCase, test := range testCases {
		input, expected := test.input, test.output
		actual := abs(input)
		if actual != expected {
			t.Errorf("Unexpected output from test %s, expected = %v, actual = %v", testCase, expected, actual)
		}
	}
}

func Test_fib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	testCases := map[string]struct {
		input  int
		output string
	}{
		"normal test 1": {
			input: 7,
			output: strconv.Quote(`1
1
2
3
5
8
13
`),
		},
		"normal test 2": {
			input: -7,
			output: strconv.Quote(`1
-1
2
-3
5
-8
13
`),
		},
		"normal test 3": {
			input:  0,
			output: strconv.Quote("0\n"),
		},
	}

	for testCase, test := range testCases {
		input, expected := test.input, test.output
		fib(input)
		actual := strconv.Quote(buf.String())
		if expected != actual {
			t.Errorf("Unexpected output from test %s, expected = %s, actual = %s", testCase, expected, actual)
		}

		buf.Reset()
	}
}

func Test_main(t *testing.T) {
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
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
