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

	expected := strconv.Quote("[1 2 3 5]\n[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

//bubble() test
func TestBubbleOutput(t *testing.T) {
	testCases := map[string]struct {
		input       []int
		expectedArr []int
	}{
		"{-20,-35}":         {input: []int{-20, -35}, expectedArr: []int{-35, -20}},
		"{0,2,0,2}":         {input: []int{0, 2, 0, 2}, expectedArr: []int{0, 0, 2, 2}},
		"{100,30000,29999}": {input: []int{100, 30000, 29999}, expectedArr: []int{100, 29999, 30000}},
	}

	for name, test := range testCases {
		test := test
		name := name
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			resultArr := bubble(test.input)
			assert.Equalf(t, resultArr, test.expectedArr, "Unexpected output in %v\nexpected: %v,\nactual: %v",
				name, test.expectedArr, resultArr)
		})
	}
}

//insertion() test
func TestInsertionOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := map[string]struct {
		input       []int
		expectedArr []int
	}{
		"{0}":          {input: []int{0}, expectedArr: []int{0}},
		"{0, 0, 0}":    {input: []int{0, 0, 0}, expectedArr: []int{0, 0, 0}},
		"{0,1,2}":      {input: []int{0, 1, 2}, expectedArr: []int{0, 1, 2}},
		"{6,3,19}":     {input: []int{6, 3, 19}, expectedArr: []int{3, 6, 19}},
		"{3, 2, 1, 5}": {input: []int{3, 2, 1, 5}, expectedArr: []int{1, 2, 3, 5}},
	}

	for name, test := range testCases {
		test := test
		name := name
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			resultArr := insertion(test.input)
			assert.Equalf(t, resultArr, test.expectedArr, "Unexpected output in %v\nexpected: %v,\nactual: %v",
				name, test.expectedArr, resultArr)
		})
	}
}
