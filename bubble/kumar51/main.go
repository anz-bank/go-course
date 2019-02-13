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

func bubble(input []int) []int {
	arrey := make([]int, len(input))
	copy(arrey, input)
	arrlen := len(arrey)
	for i := 0; i < arrlen-1; i++ {
		for j := 0; j < arrlen-i-1; j++ {
			if arrey[j] > arrey[j+1] {
				arrey[j], arrey[j+1] = arrey[j+1], arrey[j]
			}
		}
	}
	return arrey
}
