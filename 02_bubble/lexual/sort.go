package main

import (
	"fmt"
	"io"
	"os"
)

func bubble(s []int) []int {
	t := make([]int, len(s))
	copy(t, s)

	n := len(s)
	for {
		newN := 0
		for i := 1; i <= n-1; i++ {
			if t[i-1] > t[i] {
				t[i-1], t[i] = t[i], t[i-1]
				newN = i
			}
		}
		// everything is sorted after t[newN]
		n = newN
		if n <= 1 {
			break
		}
	}
	return t
}

func insertion(s []int) []int {
	t := make([]int, len(s))
	copy(t, s)

	for i := 1; i < len(t); i++ {
		for j := i; j > 0 && t[j-1] > t[j]; j-- {
			t[j-1], t[j] = t[j], t[j-1]
		}
	}
	return t
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
