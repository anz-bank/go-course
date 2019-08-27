package main

import (
	"reflect"
	"testing"

	output "github.com/joel00wood/test-helpers/capture"
)

func TestMain(t *testing.T) {
	expected := "[1 2 3 5]\n"
	actual := output.CaptureOutput(func() { main() })
	if expected != actual {
		t.Errorf("Unexpected response for input main(){bubble([]int{3, 2, 1, 5})), want=%s, got=%s",
			expected, actual)
	}
}

var testCases = map[string]struct {
	input, expected []int
}{
	"Standard input":          {[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	"Zero value":              {[]int{0}, []int{0}},
	"Trying to get coverage":  {[]int{-1, -6, -2, 0}, []int{-6, -2, -1, 0}},
	"Reverse standard input":  {[]int{5, 1, 2, 3}, []int{1, 2, 3, 5}},
	"Empty set":               {[]int{}, []int{}},
	"Pre-sorted":              {[]int{1, 2, 3, 5}, []int{1, 2, 3, 5}},
	"Reverse sorted":          {[]int{5, 3, 2, 1}, []int{1, 2, 3, 5}},
	"Single value":            {[]int{7}, []int{7}},
	"Negative standard input": {[]int{-3, -2, -1, -5}, []int{-5, -3, -2, -1}},
	"Long set": {[]int{-57, 43, 58, 77, -40, -66, 91, 83, 96,
		2, 0, -95, 25, -48, -96, -28, 2, -10, 92},
		[]int{-96, -95, -66, -57, -48, -40, -28,
			-10, 0, 2, 2, 25, 43, 58, 77, 83, 91, 92, 96}},
}

func TestBubble(t *testing.T) {
	for name, test := range testCases {
		actual := bubble(test.input)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test (%q) failed, expected=%v, got=%v",
				name, test.expected, actual)
		}
	}
}

func TestInsertion(t *testing.T) {
	for name, test := range testCases {
		actual := insertion(test.input)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test (%q) failed, expected=%v, got=%v",
				name, test.expected, actual)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for name, test := range testCases {
		actual := mergeSort(test.input)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test (%q) failed, expected=%v, got=%v",
				name, test.expected, actual)
		}
	}
}

func TestMerge(t *testing.T) {
	var testCases = map[string]struct {
		a, b, expected []int
	}{
		"Empty sets":   {[]int{}, []int{}, []int{}},
		"Half each":    {[]int{1, 2, 3}, []int{5, 7, 9}, []int{1, 2, 3, 5, 7, 9}},
		"Imbalanced a": {[]int{1, 2}, []int{3, 5, 7}, []int{1, 2, 3, 5, 7}},
		"Imbalanced b": {[]int{1, 2, 5}, []int{3, 7}, []int{1, 2, 3, 5, 7}},
		"Only a":       {[]int{1, 2, 3, 5}, []int{}, []int{1, 2, 3, 5}},
		"Only b":       {[]int{}, []int{1, 2, 3, 5}, []int{1, 2, 3, 5}},
	}
	for name, test := range testCases {
		actual := merge(test.a, test.b)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test (%q) failed, expected=%q, got=%q",
				name, test.expected, actual)
		}
	}
}
