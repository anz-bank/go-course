package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	name     string
	input    []int
	expected []int
}

type Sortfns func([]int) []int

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[1 2 3 5]\n", buf.String())
}

func TestSortFunctions(t *testing.T) {

	testCases := []Data{
		{"Empty set",
			[]int{},
			[]int{},
		},
		{"Single element set",
			[]int{10},
			[]int{10},
		},
		{"Two element set",
			[]int{20, 10},
			[]int{10, 20},
		},
		{"Worst case",
			[]int{200, 111, 108, 93, 89, 76, 68, 51, 42, 35},
			[]int{35, 42, 51, 68, 76, 89, 93, 108, 111, 200},
		},
		{"Average case a",
			[]int{6, 5, 3, 2, 8, 7, 1, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{"Average case b",
			[]int{3, 2, 1, 5},
			[]int{1, 2, 3, 5},
		},
		{"Best case",
			[]int{20, 31, 48, 53, 69, 76, 88, 91, 102, 115},
			[]int{20, 31, 48, 53, 69, 76, 88, 91, 102, 115},
		},
	}

	for i, fn := range []Sortfns{bubble, insertion} {
		fn := fn
		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name+strconv.Itoa(i), func(t *testing.T) {
				assert.Equal(t, tc.expected, fn(tc.input))
			})
		}
	}
}
