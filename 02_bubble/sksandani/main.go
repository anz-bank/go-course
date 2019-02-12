package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "Bubble: ", bubble([]int{3, 2, 1, 5}), "Insertion: ", insertion([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	n := len(s)
	for j := n; j > 0; j-- {
		for i := 1; i <= n-1; i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}
	return s
}

func insertion(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}
