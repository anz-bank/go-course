package main

import (
	"bytes"
	"testing"
)

func TestAbs(t *testing.T) {
	tesCases := map[string]struct {
		input    int
		excepted int
	}{
		"positive": {input: 1, excepted: 1},
		"negative": {-1, 1},
		"zero":     {0, 0},
	}
	for key, testCase := range tesCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			result := abs(test.input)
			if result != test.excepted {
				t.Errorf("expected %v, got %v", test.excepted, result)
			}
		})
	}
}

func TestFib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	tests := []struct {
		input    int
		excepted string
	}{
		{1, `1
`},
		{2, `1
1
`},
		{6, `1
1
2
3
5
8
`},
		{7, `1
1
2
3
5
8
13
`},
		{-7, `1
-1
2
-3
5
-8
13
`},
	}
	for _, testCase := range tests {
		var buf bytes.Buffer
		out = &buf
		fib(testCase.input)
		excepted := testCase.excepted
		actual := buf.String()
		if excepted != actual {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
		}
	}
}
