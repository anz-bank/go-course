package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	var ss []int
	for k := range s {
		ss = s[:len(s)-k-1]
		for i, val := range ss {
			if val > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
