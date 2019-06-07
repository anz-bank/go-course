package main

import (
	"bytes"
	"reflect"
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

func TestFibVal(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	type test struct {
		want  []int64
		input int
	}

	tests := []test{
		{input: -7, want: []int64{0, 1, -1, 2, -3, 5, -8, 13}},
		{input: 0, want: []int64{0, 0}},
		{input: 1, want: []int64{0, 1}},
		{input: 7, want: []int64{0, 1, 1, 2, 3, 5, 8, 13}},
		{input: -1, want: []int64{0, 1}},
	}

	for _, tc := range tests {

		got := (fib(tc.input))
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("want: %v, got: %v", tc.want, got)
		}
	}
}
