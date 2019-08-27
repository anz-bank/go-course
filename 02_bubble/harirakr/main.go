package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	n := len(s)
	cs := append([]int{}, s...)
	for i := n; i > 0; i-- {
		swapped := false
		for j := 0; j < i-1; j++ {
			if cs[j] > cs[j+1] {
				cs[j], cs[j+1] = cs[j+1], cs[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return cs
}

func insertion(s []int) []int {
	n := len(s)
	cs := append([]int{}, s...)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && cs[j-1] > cs[j]; j-- {
			cs[j], cs[j-1] = cs[j-1], cs[j]
		}
	}
	return cs
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
