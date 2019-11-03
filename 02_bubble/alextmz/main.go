package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)

	for i := len(r); i > 0; i-- {
		for j := 1; j < i; j++ {
			if r[j-1] > r[j] {
				r[j-1], r[j] = r[j], r[j-1]
			}
		}
	}
	return r
}
