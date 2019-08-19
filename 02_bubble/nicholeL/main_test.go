package main

import (
	"bytes"

	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		excepted []int
	}{
		"Happy case":                         {[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
		"slice with more identical elements": {[]int{3, 3, 1, 7, 1, 5}, []int{1, 1, 3, 3, 5, 7}},
		"single element slice":               {[]int{3}, []int{3}},
		"sorted slice":                       {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"reverted slice":                     {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"empty array":                        {[]int{}, []int{}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := bubbleSort(test.input)
			if !assert.Equal(t, test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})

	}

}

func TestInsertSort(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		excepted []int
	}{
		"Happy case":                         {[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
		"slice with more identical elements": {[]int{3, 3, 1, 7, 1, 5}, []int{1, 1, 3, 3, 5, 7}},
		"single element slice":               {[]int{3}, []int{3}},
		"sorted slice":                       {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"reverted slice":                     {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"empty array":                        {[]int{}, []int{}},
	}
	for key, testCase := range testCases {
		test := testCase
		t.Run(key, func(t *testing.T) {
			excepted := test.excepted
			actual := insertSort(test.input)
			if !assert.Equal(t, test.excepted, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
			}
		})

	}

}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	excepted := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
	if excepted != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", excepted, actual)
	}
}
