package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var sortTestsValue = []struct {
	input []int
	want  []int
}{
	{input: []int{3, 2, 1, 5}, want: []int{1, 2, 3, 5}},
	{input: []int{}, want: []int{}},
	{input: nil, want: nil},
	{input: []int{1, 0, 0, 1}, want: []int{0, 0, 1, 1}},
	{input: []int{1}, want: []int{1}},
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, test := range sortTestsValue {
		actual := bubbleSort(test.input)
		require.Equal(t, test.want, actual)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range sortTestsValue {
		actual := insertionSort(test.input)
		require.Equal(t, test.want, actual)
	}
}
