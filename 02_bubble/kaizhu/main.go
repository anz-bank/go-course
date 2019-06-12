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

func bubble(sc []int) []int {
	s := make([]int, len(sc))
	copy(s, sc)
	for i := 0; i < len(s)-1; i++ {
		flag := false
		for j := 0; j <= len(s)-2-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
			if !flag {
				return s
			}
		}
	}
	return s
}
