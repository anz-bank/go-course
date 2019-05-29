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

	// a, b := swapInt(2, 1)
	// fmt.Println(a, b)
}

func bubble(s []int) []int {

	var min = minValue(s)
	var max = maxValue(s)
	//fmt.Println("min", min, "max", max, "s", s)

	for i := 0; i < len(s); i++ {
		next := i + 1
		//fmt.Println("i", i, "next", next, "s", s)
		if next < len(s) && s[i] > s[next] {
			a, b := swapInt(s[i], s[next])
			s[i] = a
			s[next] = b
		}
	}
	if s[0] != min || s[len(s)-1] != max {
		//fmt.Println("ha ha not done yet")
		bubble(s)
	}

	return s
}

func swapInt(x, y int) (int, int) {
	return y, x
}

func minValue(s []int) (min int) {
	min = 0
	if len(s) > 0 {
		min = s[0]
	}
	for i := 0; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}
	return
}

func maxValue(s []int) (max int) {
	max = 0
	if len(s) > 0 {
		max = s[0]
	}
	for i := 0; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	return
}
