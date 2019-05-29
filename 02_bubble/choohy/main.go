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
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	//fmt.Fprintln(out, switchInt([]int{2, 1}))

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
	// 	newBubble = make([]int, s.length)

	// 	for i := 0; i < s.length; i++ {
	// 	if (s[0]>s[1])
	// 		currentBubble := switchInt([]int{2, 1})
	// 	}
	// 	fmt.Println("sum 1:", s)

	// 	// if (s[0]>s[1])
	// 	// 	s[0], s[1] = switchInt(s[0],s[1])
	// 	return s
	// 	//return int{s[1], s[0]}
	return s
}

// func switchInt(s []int) ([]int) {
// 	var a = s[0]
// 	var b = s[1]

// 	//var t []int
// 	t := []int{b, a}
// 	return t
// }

func swapInt(x, y int) (int, int) {
	return y, x
}
