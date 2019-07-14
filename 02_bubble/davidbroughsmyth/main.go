package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	cpy := make([]int, len(s))
	copy(cpy, s)
	for n := len(cpy); n > 1; n-- {
		for i := 1; i < n; i++ {
			if cpy[i-1] > cpy[i] {
				cpy[i-1], cpy[i] = cpy[i], cpy[i-1]
			}
		}
	}
	return cpy
}

func insertionSort(s []int) []int {
	cpy := make([]int, len(s))
	copy(cpy, s)
	for i, val := range s {
		j := i
		for ; j > 0 && val < cpy[j-1]; j-- {
			cpy[j] = cpy[j-1]
		}
		cpy[j] = val
	}
	return cpy
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
