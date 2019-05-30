package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

//main() test
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

//bubbleLoop() test
func TestBubbleLoopOutput(t *testing.T) {
	//test cases with descriptions.
	testCases := []struct {
		description  string
		input        []int
		expectedArr  []int
		expectedBool bool
	}{
		{description: "bubbleLoop []int{3, 2, 1, 5}", input: []int{3, 2, 1, 5},
			expectedArr: []int{2, 1, 3, 5}, expectedBool: false,
		},
		{description: "bubbleLoop []int{6,3,19}", input: []int{6, 3, 19},
			expectedArr: []int{3, 6, 19}, expectedBool: false,
		},
		{description: "bubbleLoop []int{0,1,2}", input: []int{0, 1, 2},
			expectedArr: []int{0, 1, 2}, expectedBool: true,
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			resultArr, resultBool := bubbleLoop(test.input)
			if !reflect.DeepEqual(resultArr, test.expectedArr) || resultBool != test.expectedBool {
				t.Errorf("Unexpected output in %v\nexpected: (%v,%v)\nactual: (%v,%v)",
					test.description, test.expectedArr, test.expectedBool, resultArr, resultBool)
			}
		})
	}
}
