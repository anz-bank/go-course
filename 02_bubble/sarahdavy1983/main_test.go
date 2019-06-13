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

	expected := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}

var tests = []struct {
	name  string
	input []int
	want  []int
}{
	{"valid", []int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{"same number", []int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
	{"blank", []int{}, []int{}},
	{"sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	{"negative number", []int{1, -5, 6, 9}, []int{-5, 1, 6, 9}},
}

func TestBubble(t *testing.T) {
	for _, test := range tests {
		got := (bubbleSort(test.input))
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("got %v want %v given %v", got, test.want, test.input)
		}
	}
}

func TestInsertion(t *testing.T) {
	for _, test := range tests {
		got := insertionSort(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v given %v", got, test.want, test.input)
		}
	}
}
