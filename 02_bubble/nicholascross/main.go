package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	count := len(s)
	sort := make([]int, count)
	copy(sort, s)

	for sorted := false; !sorted; {
		sorted = true

		for i := 0; i+1 < count; i++ {
			if sort[i] > sort[i+1] {
				sort[i+1], sort[i] = sort[i], sort[i+1]
				sorted = false
			}
		}
	}

	return sort
}

func insertion(s []int) []int {
	count := len(s)
	sort := make([]int, count)
	copy(sort, s)

	for i := 0; i < count; i++ {
		for j := i; j > 0 && sort[j-1] > sort[j]; j-- {
			sort[j], sort[j-1] = sort[j-1], sort[j]
		}
	}

	return sort
}
