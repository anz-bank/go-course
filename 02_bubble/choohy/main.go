package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "Bubble Sort!")
	//fmt.Fprintln(out, "[1 2]")

	//fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	// fmt.Fprintln(out, bubble([]int{2, 1}))
	fmt.Fprintln(out, switchInt([]int{2, 1}))
}

func bubble(s []int) []int {
	for 
	if (s[0]>s[1])
		s[0], s[1] = switchInt(s[0],s[1])
	return s
	//return int{s[1], s[0]}
}

func switchInt(s []int) []int {
	var a = s[0]
	var b = s[1]

	//var t []int
	t := []int{b, a}
	return t
}
