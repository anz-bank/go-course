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

	expected := strconv.Quote(`[1 2 3 5]
[1 2 3 5]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main() %v", actual)
	}
}

var testCases = map[string]struct {
	input []int
	want  []int
}{
	"sample1": {input: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
	"sample2": {input: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	"sample3": {input: []int{1, 2, 4, 3, 5}, want: []int{1, 2, 3, 4, 5}},
	"sample4": {input: []int{4, 3, 1, 5, 2}, want: []int{1, 2, 3, 4, 5}},
	"sample5": {input: []int{5, 1, 2, 4, 3}, want: []int{1, 2, 3, 4, 5}},
	"sample6": {input: []int{-10, 2, -5, 1, 2, 44, 3}, want: []int{-10, -5, 1, 2, 2, 3, 44}},
}

func TestBubbleSort(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, bubble(input), want, "%v was not sorted", input)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for name, test := range testCases {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, insertion(input), want, "%v was not sorted", input)
		})
	}
}
