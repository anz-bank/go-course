package main

import (
	"bytes"
	"reflect"
	"testing"
)

var testData = map[string]struct{ input, expected []int }{
	"Empty Array":        {[]int{}, []int{}},
	"Single Digit Array": {[]int{1}, []int{1}},
	"Standard Input":     {[]int{5, 8, 7}, []int{5, 7, 8}},
	"Negative Input":     {[]int{-4, -6, -2}, []int{-6, -4, -2}},
	"Repetitive Set":     {[]int{5, 6, 8, 5, 3, 8, 7, 6}, []int{3, 5, 5, 6, 6, 7, 8, 8}},
	"Long Set":           {[]int{12, 15, 8, 5, 3, 2, 4, 1, 0}, []int{0, 1, 2, 3, 4, 5, 8, 12, 15}},
}

func TestBubbleSort(t *testing.T) {
	for k, v := range testData {
		input := append([]int{}, v.input...)
		actual := bubbleSort(v.input)
		if !reflect.DeepEqual(input, v.input) {
			t.Errorf("InsertionSort Original Input %v Modified to %v ", input, v.input)
		}
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("BubbleSort Test Case %v Expected %v,Actual %v", k, v.expected, actual)
		}
	}
}
func TestInsertionSort(t *testing.T) {
	for k, v := range testData {
		input := append([]int{}, v.input...)
		actual := insertionSort(v.input)
		if !reflect.DeepEqual(input, v.input) {
			t.Errorf("InsertionSort Original Input %v Modified to %v ", input, v.input)
		}
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("InsertionSort Test Case %v Expected %v,Actual %v", k, v.expected, actual)
		}
	}
}
func TestMain(t *testing.T) {
	var b bytes.Buffer
	out = &b
	main()
	expected := "[1 2 3 5]\n"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected %v and Actual %v  Not Matching", expected, actual)
	}
}
