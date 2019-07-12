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

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected %v, sadly got %v", expected, actual)
	}
}

var testCases = map[string]struct {
	input []int
	want  []int
}{
	"empty":           {input: []int{}, want: []int{}},
	"single value":    {input: []int{0}, want: []int{0}},
	"positive values": {input: []int{1, 6, 2, 0}, want: []int{0, 1, 2, 6}},
	"negative values": {input: []int{-1, -6, -2, 0}, want: []int{-6, -2, -1, 0}},
	"already sorted":  {input: []int{0, 1, 2, 6}, want: []int{0, 1, 2, 6}},
	"reverse sorted":  {input: []int{6, 2, 1, 0}, want: []int{0, 1, 2, 6}},
}

func TestBubble(t *testing.T) {
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			b := bubble(test.input)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("expected %v, got %v", test.want, b)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			b := insertion(test.input)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("expected %v, got %v", test.want, b)
			}
		})
	}
}

func TestMergesort(t *testing.T) {
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			b := mergesort(test.input)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("expected %v, got %v", test.want, b)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	var testCases = map[string]struct {
		left  []int
		right []int
		want  []int
	}{
		"empty":         {left: []int{}, right: []int{}, want: []int{}},
		"less in left":  {left: []int{4, 0}, right: []int{6, 2, 4}, want: []int{4, 0, 6, 2, 4}},
		"less in right": {left: []int{2, 4, 0}, right: []int{6, 4}, want: []int{2, 4, 0, 6, 4}},
		"even":          {left: []int{2, 4, 6}, right: []int{7, 5, 3}, want: []int{2, 4, 6, 7, 5, 3}},
		"left only":     {left: []int{1, 6, 2, 0}, right: []int{}, want: []int{1, 6, 2, 0}},
		"right only":    {left: []int{}, right: []int{1, 6, 2, 0}, want: []int{1, 6, 2, 0}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			b := merge(test.left, test.right)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("expected %v, got %v", test.want, b)
			}
		})
	}
}
