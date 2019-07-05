package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(format("1 1 2 3 5 8 13"))
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("%s does not match expected %s", actual, expected)
	}
}

func TestFib(t *testing.T) {
	cases := []struct {
		length   int
		sequence string
	}{
		{length: 7, sequence: "1 1 2 3 5 8 13"},
		{length: -7, sequence: "1 -1 2 -3 5 -8 13"},
		{length: 0, sequence: ""},
		{length: 1, sequence: "1"},
		{length: 10, sequence: "1 1 2 3 5 8 13 21 34 55"},
	}

	for _, test := range cases {
		var buf bytes.Buffer
		out = &buf
		length := test.length
		want := format(test.sequence)

		t.Run(string(length), func(t *testing.T) {
			fib(length)
			assert.Equalf(t, buf.String(), want, "%v was incorrect", length)
		})
	}
}

func format(sequence string) string {
	return strings.Replace(sequence, " ", "\n", -1)
}
