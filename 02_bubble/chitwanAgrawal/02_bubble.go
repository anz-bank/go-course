package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
