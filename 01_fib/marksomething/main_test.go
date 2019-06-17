package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibString(t *testing.T) {
	testCases := map[string]struct {
		value    int
		expected string
	}{
		"fib(0)": {
			value:    0,
			expected: "0\n",
		},
		"fib(1)": {
			value:    1,
			expected: "0\n1\n",
		},
		"fib(2)": {
			value:    2,
			expected: "0\n1\n1\n",
		},
		"fib(10)": {
			value:    10,
			expected: "0\n1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n",
		},
		"fib(-1)": {
			value:    -1,
			expected: "0\n1\n",
		},
		"fib(-2)": {
			value:    -2,
			expected: "0\n1\n-1\n",
		},
		"fib(-10)": {
			value:    -10,
			expected: "0\n1\n-1\n2\n-3\n5\n-8\n13\n-21\n34\n-55\n",
		},
	}
	for testName, tC := range testCases {
		value := tC.value
		expected := tC.expected
		t.Run(testName, func(t *testing.T) {
			actual := fibString(value)
			expected := expected
			assert.Equal(t, expected, actual)
		})
	}
}

func TestFib(t *testing.T) {
	testCases := map[string]struct {
		value    int
		expected string
	}{
		"fib(0)": {
			value: 0,
			expected: `0
`,
		},
		"fib(-1)": {
			value: -1,
			expected: `0
1
`,
		},
		"fib(1)": {
			value: 1,
			expected: `0
1
`,
		},
		"fib(10)": {
			value: 10,
			expected: `0
1
1
2
3
5
8
13
21
34
55
`,
		},
		"fib(-10)": {
			value: -10,
			expected: `0
1
-1
2
-3
5
-8
13
-21
34
-55
`,
		},
	}
	for testName, tC := range testCases {
		value := tC.value
		expected := tC.expected
		t.Run(testName, func(t *testing.T) {
			var buf bytes.Buffer
			out = &buf
			fib(value)
			actual := buf.String()
			assert.Equal(t, expected, actual)

		})
	}
}

func TestExtremePositive(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	fibString(93)
}

func TestExtremeNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	fibString(-93)
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()
	expected := `
Fib 7
------
0
1
1
2
3
5
8
13
`
	if actual != expected {
		t.Errorf("\nActual:\n%v\nExpected:\n%v", actual, expected)
	}

}
