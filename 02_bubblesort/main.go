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
	n := len(s) - 1
	for {
		if n <= 0 {
			break
		}

		for i := 0; i < n; i++ {
			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
			}
		}
		n--
	}
	return sorted
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
