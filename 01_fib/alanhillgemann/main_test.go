package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
}

var testCases = map[string]struct {
	input    int
	expected string
}{
	"Neg": {input: -5, expected: ""},
	"0":   {input: 0, expected: ""},
	"1":   {input: 1, expected: "1\n"},
	"2":   {input: 2, expected: "1\n1\n"},
	"7":   {input: 7, expected: "1\n1\n2\n3\n5\n8\n13\n"},
	"45 max": {input: 45, expected: "1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n89\n144\n233\n377\n610\n" +
		"987\n1597\n2584\n4181\n6765\n10946\n17711\n28657\n46368\n75025\n121393\n196418\n" +
		"317811\n514229\n832040\n1346269\n2178309\n3524578\n5702887\n9227465\n14930352\n" +
		"24157817\n39088169\n63245986\n102334155\n165580141\n267914296\n433494437\n" +
		"701408733\n1134903170\n"},
	"46 to 45": {input: 46, expected: "1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n89\n144\n233\n377\n610\n" +
		"987\n1597\n2584\n4181\n6765\n10946\n17711\n28657\n46368\n75025\n121393\n196418\n" +
		"317811\n514229\n832040\n1346269\n2178309\n3524578\n5702887\n9227465\n14930352\n" +
		"24157817\n39088169\n63245986\n102334155\n165580141\n267914296\n433494437\n" +
		"701408733\n1134903170\n"},
}

func TestFib(t *testing.T) {
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			out = &buf
			fib(test.input)
			actual := strconv.Quote(buf.String())
			expected := strconv.Quote(test.expected)
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}
