package main

import (
	"fmt"
	"os"
	"io"
)

func bubble(s []int) []int {
	/*
		Returns a sorted copy of slice s using bubble sort
	*/

	// Edge case: empty or length 1
	if len(s) < 2 {
		return s
	}
	// Create a ptr from 0 to length - 1
	repeat := true
	for repeat {
		repeat = false
		for ptr := 0; ptr < len(s)-1; ptr++ {
			// Compare the current index with the next index, if differnet then exchange
			if s[ptr] > s[ptr+1] {
				repeat = true
				s[ptr], s[ptr+1] = s[ptr+1], s[ptr]
			}
		}
	}
	return s
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3,2,1,5}))
}