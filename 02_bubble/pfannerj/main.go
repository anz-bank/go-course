package main

import (
	"fmt"
	"io"
	"os"
)

var mainout io.Writer = os.Stdout

func main() {
	fmt.Fprintf(mainout, "Bubble sort starting...\n")
	fmt.Fprintln(mainout, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintf(mainout, "Bubble sort completed...\n")

	fmt.Fprintf(mainout, "Insertion sort starting...\n")
	fmt.Fprintln(mainout, insertion([]int{3, 2, 5, 1}))
	fmt.Fprintf(mainout, "Insertion sort completed...\n")
}

func bubble(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)
	var saveNbr int
	swapped := false
	i := 0
	for {
		if (i + 1) >= len(s) {
			i = 0
			if swapped {
				swapped = false
			} else {
				break //Sort is complete if no values were 'swapped' during the previous iteration
			}
		}
		if sCopy[i] > sCopy[i+1] {
			saveNbr = sCopy[i]
			sCopy[i], sCopy[i+1] = sCopy[i+1], saveNbr
			swapped = true
		}
		i++
	}
	return sCopy
}

func insertion(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)
	var moveNbr, saveNbr int
	for i := 0; i < len(sCopy)-1; i++ {
		if sCopy[i+1] < sCopy[i] {
			moveNbr = sCopy[i+1]      //This value is out of order and needs to be moved
			for j := i; j >= 0; j-- { //Loop to place the moveNbr in correct location
				if moveNbr < sCopy[j] {
					saveNbr = sCopy[j]
					sCopy[j], sCopy[j+1] = moveNbr, saveNbr
				}
			}
		}
	}
	return sCopy
}
