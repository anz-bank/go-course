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
	sCopy := make([]int, len(s))
	copy(sCopy, s)

	for i := len(s); i > 0; i-- {
		swapped := false

		for j := 1; j < i; j++ {
			if sCopy[j-1] > sCopy[j] {
				sCopy[j-1], sCopy[j] = sCopy[j], sCopy[j-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return sCopy
}
