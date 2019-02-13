package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(items []int) []int {
	n := len(items) - 1
	for {
		if n <= 0 {
			break
		}

		for i := 0; i < n; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
			}
		}
		n--
	}
	return items
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
