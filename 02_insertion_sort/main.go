package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func insertion(items []int) []int {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}
	return items
}

func main() {
	fmt.Fprint(out, insertion([]int{3, 2, 1, 5}))
}
