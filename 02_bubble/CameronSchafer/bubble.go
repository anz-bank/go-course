package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	newS, check := s, false

	//loop until check becomes true
	for !check {
		//sort once
		newS, check = bubbleLoop(newS)
	}

	return s
}

//this function loops through the slice and bubble sorts it
func bubbleLoop(nextS []int) ([]int, bool) {
	bubbleCheck := true

	//loop through the slice
	for i := 0; i < len(nextS); i++ {
		//check if next slice elem exists
		if i+1 < len(nextS) {
			//check if current elem is bigger than the next elem
			if nextS[i] > nextS[i+1] {
				//swap them around
				nextS[i], nextS[i+1] = nextS[i+1], nextS[i]
				bubbleCheck = false
			}
		} else {
			//break at last slice elem
			break
		}
	}

	return nextS, bubbleCheck
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
