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
	n := len(s)
	k := make([]int, n)
	copy(k, s)
	for j := n; j > 0; j-- {
		for i := 1; i <= n-1; i++ {
			if k[i-1] > k[i] {
				k[i], k[i-1] = k[i-1], k[i]
			}
		}
	}
	return k
}

func insertion(s []int) []int {
	n := len(s)
	k := make([]int, n)
	copy(k, s)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if k[j-1] > k[j] {
				k[j], k[j-1] = k[j-1], k[j]
			}
		}
	}
	return k
}
