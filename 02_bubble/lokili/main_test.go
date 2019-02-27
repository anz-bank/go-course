package main

import (
	"bytes"
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
	expected := `[1 2 3 5]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var testSet = map[string]map[string][]int{
	"testSet1": {
		"input":  {-10, -110, 100, 0, 4, 60, 43},
		"output": {-110, -10, 0, 4, 43, 60, 100},
	},
	"testSet2": {
		"input":  {4, 2, 0},
		"output": {0, 2, 4},
	},
	"testSet3": {
		"input":  {1, 2, 3},
		"output": {1, 2, 3},
	},
}

func TestBubbleSort(t *testing.T) {
	r := require.New(t)

	for _, test := range testSet {
		sorted := bubble(test["input"])
		r.Equalf(test["output"], sorted, "Unexpected output in bubble()")
	}
}

func TestInsertionSort(t *testing.T) {
	r := require.New(t)

	for _, test := range testSet {
		sorted := insertion(test["input"])
		r.Equalf(test["output"], sorted, "Unexpected output in insertion()")
	}
}

func TestMergeSort(t *testing.T) {
	r := require.New(t)

	for _, test := range testSet {
		sorted := merge(test["input"])
		r.Equalf(test["output"], sorted, "Unexpected output in merge()")
	}
}
