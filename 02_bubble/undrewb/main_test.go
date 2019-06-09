package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("[1 2 3 5]\n")
	got := strconv.Quote(buf.String())

	if got != want {
		t.Errorf("actual: %s does not match expected: %s", got, want)
	}
}

var sortTests = []struct {
	name  string
	input []int
	want  []int
}{
	{name: "lab example", input: []int{3, 2, 1, 5}, want: []int{1, 2, 3, 5}},
	{name: "odd entries", input: []int{3, 2, 1, 6, 5}, want: []int{1, 2, 3, 5, 6}},
	{name: "duplicate entries", input: []int{3, 4, 2, 1, 4, 5}, want: []int{1, 2, 3, 4, 4, 5}},
	{name: "empty", input: []int{}, want: []int{}},
	{name: "one entry", input: []int{1}, want: []int{1}},
	{name: "repeated entry", input: []int{3, 3, 3}, want: []int{3, 3, 3}},
	{name: "already sorted", input: []int{1, 2, 3, 5}, want: []int{1, 2, 3, 5}},
	{name: "reverse sorted", input: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	{name: "long example", input: []int{3, 2, 1, 5, 7, 8, 6, 4, 2, 3, 1},
		want: []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7, 8}},
}

func TestBubble(t *testing.T) {
	for _, tt := range sortTests {
		tt := tt // as per sean- suggestion on this discussion https://github.com/kyoh86/scopelint/issues/4
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	for _, tt := range sortTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := insertion(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	for _, tt := range sortTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

var mergeTests = []struct {
	name  string
	left  []int
	right []int
	want  []int
}{
	{name: "simple example", left: []int{3, 5}, right: []int{2, 1}, want: []int{2, 1, 3, 5}},
	{name: "simple example", left: []int{2, 1}, right: []int{3, 5}, want: []int{2, 1, 3, 5}},
	{name: "empty example", left: []int{}, right: []int{}, want: []int{}},
	{name: "empty left", left: []int{}, right: []int{2, 1}, want: []int{2, 1}},
	{name: "empty right", left: []int{3, 5}, right: []int{}, want: []int{3, 5}},
	{name: "lopsided left", left: []int{3, 5, 7}, right: []int{2, 1}, want: []int{2, 1, 3, 5, 7}},
	{name: "lopsided right", left: []int{3, 5}, right: []int{1, 8, 2}, want: []int{1, 3, 5, 8, 2}},
}

func TestMergeSlice(t *testing.T) {
	for _, tt := range mergeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSlice(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
