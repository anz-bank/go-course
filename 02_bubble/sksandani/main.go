package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
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
