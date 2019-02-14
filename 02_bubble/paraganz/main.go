package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(arr []int) []int {
	s := make([]int, len(arr))
	copy(s, arr)
	cnt := len(s)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
				j--
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
