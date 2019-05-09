package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	l := len(s) - 1
	fin := false
	for !fin {
		sw := false
		for i := 0; i < l; i++ {
			if s[i] > s[i+1] {
				sw = true
				s[i], s[i+1] = s[i+1], s[i]
			}
			fin = !sw
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
