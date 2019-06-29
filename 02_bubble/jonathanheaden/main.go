package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	done := false
	if len(s) == 0 {
		return s
	}
	r := make([]int, len(s))
	copy(r, s)
	for !done {
		sw := false
		for i := 0; i < len(r)-1; i++ {
			if r[i] > r[i+1] {
				sw = true
				r[i], r[i+1] = r[i+1], r[i]
			}
			done = !sw
		}
	}
	return r
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
