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
	cs1 := make([]int, len(s))
	copy(cs1, s)
	end := len(cs1) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(cs1)-1; i++ {
			if cs1[i] > cs1[i+1] {
				cs1[i], cs1[i+1] = cs1[i+1], cs1[i]
			}
		}
		end--
	}
	return cs1
}

func insertionsort(items []int) []int {
	cs2 := make([]int, len(items))
	copy(cs2, items)
	var n = len(cs2)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if cs2[j-1] > cs2[j] {
				cs2[j-1], cs2[j] = cs2[j], cs2[j-1]
			}
			j = j - 1
		}
	}
	return cs2
}
