package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf
	main()
	expected := `Bubble sort starting...
[1 2 3 5]
Bubble sort completed...
Insertion sort starting...
[1 2 3 5]
Insertion sort completed...
`
	actual := buf.String()
	if expected != actual {
		t.Errorf(expected, actual, "Unexpected output in main()")
	}
}

var testCases = map[string]struct {
	input    []int
	expected []int
}{
	"Four":        {input: []int{3, 1, 2, 5}, expected: []int{1, 2, 3, 5}},
	"Eight":       {input: []int{6, 3, 8, 7, 5, 4, 1, 2}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8}},
	"RepeatedNbr": {input: []int{6, 3, 1, 7, 5, 4, 1, 2}, expected: []int{1, 1, 2, 3, 4, 5, 6, 7}},
	"AllSame":     {input: []int{6, 6, 6, 6, 6, 6, 6, 6}, expected: []int{6, 6, 6, 6, 6, 6, 6, 6}},
	"NegativeNbr": {input: []int{3, -2, 1, 5}, expected: []int{-2, 1, 3, 5}},
	"Empty":       {input: []int{}, expected: []int{}},
}

func TestBubbleSort(t *testing.T) {
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := bubble(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Bubble sort failed. Input: %v, Actual: %v", tc.input, actual)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := insertion(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Insertion sort failed. Input: %v, Actual: %v", tc.input, actual)
			}
		})
	}
}
