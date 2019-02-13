package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func insertion(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	var n = len(s)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if sorted[j-1] > sorted[j] {
				sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
			}
			j--
		}
	}
	return sorted
}

func main() {
	fmt.Fprint(out, insertion([]int{3, 2, 1, 5}))
}
