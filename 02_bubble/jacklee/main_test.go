package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	//when
	main()

	//Then
	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	r.JSONEqf(expected, actual, "Unexpected output in main()")
}

var testSet = map[string]map[string][]int{
	"testSet1": {
		"input":  {3, 2, 1, -2, 10, 19, 0, 81, -29},
		"output": {-29, -2, 0, 1, 2, 3, 10, 19, 81},
	},
	"testSet2": {
		"input":  {3, 2, 1},
		"output": {1, 2, 3},
	},
	"testSet3": {
		"input":  {20, 40, 80, 100},
		"output": {20, 40, 80, 100},
	},
	"testSet4": {
		"input":  {-3, 20, 19210, 110},
		"output": {-3, 20, 110, 19210},
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

func TestHeapSort(t *testing.T) {
	//Given
	r := require.New(t)

	for _, test := range testSet {
		res := heapSort(test["input"])
		r.Equalf(test["output"], res, "Unexpected output in heapSort()")
	}
}
