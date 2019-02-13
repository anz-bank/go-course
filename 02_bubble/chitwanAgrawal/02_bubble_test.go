package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputs = []struct {
	in []int
	out []int
}{
	{[]int{0, 9, 8, 3, 6, 2}, []int{0, 2, 3, 6, 8, 9}},
	{[]int{1, 2, 3, 4, 5, 10, 12}, []int{1, 2, 3, 4, 5, 10, 12}},
	{[]int{12, 10, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 10, 12}},
	{[]int{12, 10, 5, 4, 3, 5, 2, 1, 12, 10}, []int{1, 2, 3, 4, 5, 5, 10, 10, 12, 12}},
	{[]int{-9, -10, -3, -4, -6, -2}, []int{-10, -9, -6, -4, -3, -2}}}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[1 2 3 5]\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestBubbleSortOutput(t *testing.T) {
   r := require.New(t)	
   for _, t := range testInputs {
	   actual := bubble(t.in)
       r.ElementsMatch(t.out, actual)		
   } 
}

