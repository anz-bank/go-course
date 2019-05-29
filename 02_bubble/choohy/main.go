package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "Bubble Sort!")

	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))

	a, b := swapInt(2, 1)
	fmt.Println(a, b)
}

func bubble(s []int) []int {

	//fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		next := i + 1
		fmt.Println("i", i, "next", next, "s", s)
		if next < len(s) && s[i] > s[next] {
			a, b := swapInt(s[i], s[next])
			s[i] = a
			s[next] = b
		}
	}
	return s
}

func swapInt(x, y int) (int, int) {
	return y, x
}
