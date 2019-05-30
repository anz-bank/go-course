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

	expected := strconv.Quote("[3 2 1 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
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
		{description: "merge []int{3, 2, 1, 5}", input: []int{3, 2, 1, 5},
			expectedArr: []int{3, 2, 1, 5},
		},
		{description: "merge []int{6,3,19}", input: []int{6, 3, 19},
			expectedArr: []int{6, 3, 19},
		},
		{description: "merge []int{0,1,2}", input: []int{0, 1, 2},
			expectedArr: []int{0, 1, 2},
		},
	}

	for _, test := range testCases {
		test := test
		// t.Run creates a sub test and runs it like a normal test
		t.Run(test.description, func(t *testing.T) {
			resultArr := merge(test.input)
			if !reflect.DeepEqual(resultArr, test.expectedArr) {
				t.Errorf("Unexpected output in %v\nexpected: %v,\nactual: %v",
					test.description, test.expectedArr, resultArr)
			}
		})
	}
}
