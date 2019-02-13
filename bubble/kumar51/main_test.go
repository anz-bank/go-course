package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var testSet = map[string]map[string][]int{
	"testSet1": {
		"input":  {4, 2, 1, -2, 10, 19, 0, 81, -20},
		"output": {-20, -2, 0, 1, 2, 4, 10, 19, 81},
	},
	"testSet2": {
		"input":  {1, 8, 4},
		"output": {1, 4, 8},
	},
	"testSet3": {
		"input":  {1, 2, 3, 4},
		"output": {1, 2, 3, 4},
	},
	"testSet4": {
		"input":  {-7, 20, 19210, 110},
		"output": {-7, 20, 110, 19210},
	},
}

func TestBubbleSort(t *testing.T) {
	// Given
	r := require.New(t)

	for _, test := range testSet {
		res := bubble(test["input"])
		r.Equalf(test["output"], res, "Unexpected output in bubble()")
	}
}
