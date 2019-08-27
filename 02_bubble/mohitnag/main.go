package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	for j := 0; j < len(s); j++ {
		swapped := false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				swap(s, i-1, i)
				swapped = true
			}
		}
		if !swapped {
			return s
		}
	}
	return s
}
func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
