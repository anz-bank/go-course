package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
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
	expected := strconv.Quote(`[-2092 -10 -5 -2 -2 0 2 3 9 19 29][-2092 -10 -5 -2 -2 0 2 3 9 19 29]`)
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleAndInsertion(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	testSet := map[string]map[string][]int{
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
		"testSet5": {
			"input":  {3, 2, 1, -2, 10, 19, 0, 81, -29},
			"output": {-29, -2, 0, 1, 2, 3, 10, 19, 81},
		},
		"testSet6": {
			"input":  {3, 2, 1},
			"output": {1, 2, 3},
		},
		"testSet7": {
			"input":  {20, 40, 80, 100},
			"output": {20, 40, 80, 100},
		},
		"testSet8": {
			"input":  {-3, 20, 19210, 110},
			"output": {-3, 20, 110, 19210},
		},
	}

	for k, test := range testSet {
		var pass bool
		if strings.HasPrefix(k, "testSet") {
			res := bubble(test["input"])
			if !reflect.DeepEqual(res, test["output"]) {
				pass = false
			}
		}
		r.Equalf(pass, false, "Unexpected output in main()")
	}
}
