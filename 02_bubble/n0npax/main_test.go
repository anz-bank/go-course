package main

import (
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}

func randomSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(100) - 50
	}
	return s
}

type fn func([]int) []int

func TestSort(t *testing.T) {
	testCase := map[string]fn{
		"bubble":    bubble,
		"insertion": insertion,
		"quicksort": quicksort,
	}
	for size := 0; size <= 10; size++ {
		for name, testFunc := range testCase {
			testFunc, size := testFunc, size
			t.Run(name, func(t *testing.T) {
				s := randomSlice(size)
				expected := append(s[:0:0], s...)
				sort.Ints(expected)
				actual := testFunc(s)
				if expected == nil || actual == nil {
					t.Errorf("Unexpected nil. Shouldn't happen")
				}
				if len(expected) != len(actual) {
					t.Errorf("Unexpected slice len. Expected: %d - Actual: %d", expected, actual)
				}
				for j := 0; j < size; j++ {
					if expected[j] != actual[j] {
						t.Errorf("Index: %d - Expected: %d - Actual: %d", j, expected[j], actual[j])
						t.Error("Expected: ", expected)
						t.Error("Actual: ", actual)
					}
				}

			})
		}
	}
}
