package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_bubble(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          []int
		output         []int
	}{
		{"normal test 1", []int{1, 6, 4, 7, 5, 3, 2}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"normal test 2", []int{6, -32, 0, 0, 57, 1, 7}, []int{-32, 0, 0, 1, 6, 7, 57}},
		{"empty input", []int{}, []int{}},
	}
	for _, table := range tables {
		actual := bubble(table.input)
		expected := table.output
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected output from test %s, expected = %v, actual = %v", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_insertion(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          []int
		output         []int
	}{
		{"normal test 1", []int{1, 6, 4, 7, 5, 3, 2}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"normal test 2", []int{6, -32, 0, 0, 57, 1, 7}, []int{-32, 0, 0, 1, 6, 7, 57}},
		{"empty input", []int{}, []int{}},
	}
	for _, table := range tables {
		actual := insertion(table.input)
		expected := table.output
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected output from test %s, expected = %v, actual = %v", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_quicksort(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          []int
		output         []int
	}{
		{"normal test 1", []int{1, 6, 4, 7, 5, 3, 2}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"normal test 2", []int{6, -32, 0, 0, 57, 1, 7}, []int{-32, 0, 0, 1, 6, 7, 57}},
		{"empty input", []int{}, []int{}},
	}
	for _, table := range tables {
		actual := quicksort(table.input)
		expected := table.output
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected output from test %s, expected = %v, actual = %v", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_main(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "[0 1 2 2 3 5 6 8]"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}
