package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	n := len(s)
	for {
		if n == 0 {
			break
		}
		for i := 0; i < n-1; i++ {
			if sorted[i] > sorted[i+1] {
				sorted[i+1], sorted[i] = sorted[i], sorted[i+1]
			}
		}
		n--
	}
	return sorted
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
