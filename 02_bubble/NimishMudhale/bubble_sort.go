package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	n := len(s)
	for {
		if n == 0 {
			break
		}
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
		n--
	}
	return s
}
func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
