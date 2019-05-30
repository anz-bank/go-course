package main

import (
	"bytes"
	"strconv"
	"testing"
)

//main() test
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
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestFibOutput(t *testing.T) {
	//tempBuf is used to reset the output buffer for each test
	var buf, tempBuf bytes.Buffer
	out = &buf

	//test cases with descriptions.
	testCases := []struct {
		description string
		input       int
		expected    string
	}{
		{description: "fib 7", input: 7,
			expected: strconv.Quote(`1
1
2
3
5
8
13
`),
		},
		{description: "fib 1", input: 1,
			expected: strconv.Quote("1\n"),
		},
		{description: "fib -7", input: -7,
			expected: strconv.Quote(`1
-1
2
-3
5
-8
13
`),
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			fib(test.input)
			result := strconv.Quote(buf.String())
			if result != test.expected {
				t.Errorf("%v\nexpected %v, got %v", test.description, test.expected, result)
			}
			buf = tempBuf //reset the buffer for the next test.
		})
	}
}
