package main

import (
	"fmt"
	"io"
	"os"
)

func bubble(s []int) []int {
	n := len(s)
	var copiedSlice = make([]int, n)
	copy(copiedSlice, s)
	var swapped bool
	for ok := true; ok; ok = swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if copiedSlice[i-1] > copiedSlice[i] {
				copiedSlice[i-1], copiedSlice[i] = copiedSlice[i], copiedSlice[i-1]
				swapped = true
			}
		}
	}
	return copiedSlice
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))

}
