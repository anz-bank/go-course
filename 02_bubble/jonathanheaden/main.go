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
	for !done {
		sw := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				sw = true
				s[i], s[i+1] = s[i+1], s[i]
			}
			done = !sw
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
