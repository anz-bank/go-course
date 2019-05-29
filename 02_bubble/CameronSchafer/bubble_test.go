package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//main() test
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

//bubble() test
func TestBubbleOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description string
		input       []int
		expectedArr []int
	}{
		{description: "bubble []int{3, 2, 1, 5}", input: []int{3, 2, 1, 5},
			expectedArr: []int{1, 2, 3, 5},
		},
		{description: "bubble []int{6,3,19}", input: []int{7, -1, 10000},
			expectedArr: []int{-1, 7, 10000},
		},
		{description: "bubble []int{0,1,2}", input: []int{-1, -30, -2},
			expectedArr: []int{-30, -2, -1},
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			resultArr := bubble(test.input)
			assert.Equalf(t, resultArr, test.expectedArr, "Unexpected output in %v\nexpected: %v,\nactual: %v",
				test.description, test.expectedArr, resultArr)
		})
	}
}

//insertion() test
func TestInsertionOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description string
		input       []int
		expectedArr []int
	}{
		{description: "insertion []int{3, 2, 1, 5}", input: []int{3, 83, 103003, 50, 20, 29, 5218, -1295, 0},
			expectedArr: []int{-1295, 0, 3, 20, 29, 50, 83, 5218, 103003},
		},
		{description: "insertion []int{6,3,19}", input: []int{6},
			expectedArr: []int{6},
		},
		{description: "insertion []int{0,1,2}", input: []int{0, 1},
			expectedArr: []int{0, 1},
		},
		{description: "insertion []int{0,1,2}", input: []int{-10, 1, -12, 2, 90, 0, -100},
			expectedArr: []int{-100, -12, -10, 0, 1, 2, 90},
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			resultArr := insertion(test.input)
			assert.Equalf(t, resultArr, test.expectedArr, "Unexpected output in %v\nexpected: %v,\nactual: %v",
				test.description, test.expectedArr, resultArr)
		})
	}
}

//merge() test
func TestMergeOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description string
		input       []int
		expectedArr []int
	}{
		{description: "merge []int{3, 2, 1, 5}", input: []int{-3, -2, -1, -5},
			expectedArr: []int{-5, -3, -2, -1},
		},
		{description: "merge []int{0,1,2}", input: []int{-10, 1, -12, 2, 90, 0, -100, -10, 1, -12, 2, 90, 0, -100},
			expectedArr: []int{-100, -100, -12, -12, -10, -10, 0, 0, 1, 1, 2, 2, 90, 90},
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			resultArr := merge(test.input)
			assert.Equalf(t, resultArr, test.expectedArr, "Unexpected output in %v\nexpected: %v,\nactual: %v",
				test.description, test.expectedArr, resultArr)
		})
	}
}
