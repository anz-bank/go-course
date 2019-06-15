package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func makeCopy(ns []int) []int {
	return append([]int{}, ns...)
}

func bubble(s []int) []int {
	copy := makeCopy(s)
	n := len(copy)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			maybeSwap(copy, j, j+1)
		}
	}
	return copy
}

func maybeSwap(ns []int, i, j int) bool {
	if ns[i] > ns[j] {
		ns[i], ns[j] = ns[j], ns[i]
		return true
	}
	return false
}

func insertion(ns []int) []int {
	n := len(ns)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			swapped := maybeSwap(ns, j, j+1)
			if !swapped {
				break
			}
		}
	}
	return ns
}

func main() {
	ns := []int{3, 2, 1}
	fmt.Fprintln(out, bubble(ns))
	fmt.Fprintln(out, insertion(ns))
}
