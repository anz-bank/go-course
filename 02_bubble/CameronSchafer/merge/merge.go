package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//merge sort: O(n log n) sorting algorithm
func merge(s []int) []int {
	var sliceContainer [][]int
	check := false

	//split the orginal slice into slices of size 1
	for len(s) > 0 {
		sliceContainer = append(sliceContainer, s[:1])
		s = s[1:]
	}
	fmt.Fprintln(out, sliceContainer)

	var tempSliceContainer [][]int
	var tempSlice []int
	for !check {
		if(len(tempSliceContainer) > 0){
			sliceContainer = tempSliceContainer
			if(len(sliceContainer) == 1){
				check = true
			}
		}
		for i := 0; i < len(sliceContainer); i++ {
			//clear old temp slice
			tempSlice = []int{}
			//current slice
			for j := 0; j < len(sliceContainer[i]); j++{
				//fmt.Fprintln(out, sliceContainer[i])
				if(i + 1 < len(sliceContainer)){
					//adjacent slice
					for k := 0; k < len(sliceContainer[i + 1]); k++ {
						//fmt.Fprintln(out, sliceContainer[i + 1])
						//compare the pair
						if(sliceContainer[i][j] < sliceContainer[i + 1][k]){
							tempSlice = append(tempSlice, sliceContainer[i][j])
							sliceContainer[i] = sliceContainer[i][1:]
						}else{
							tempSlice = append(tempSlice, sliceContainer[i + 1][k])
							sliceContainer[i] = sliceContainer[i + 1][1:]
						}
					}
				}else{
					fmt.Fprintln(out, "no more elements in the slice")
				}
			}
			//put new tempslice into the tempslice container
			
			tempSliceContainer = append(tempSliceContainer, tempSlice)
			check = true
		}
	}
	fmt.Fprintln(out, tempSliceContainer)
	return s
}

func main() {
	fmt.Fprintln(out, merge([]int{3, 2, 1, 5, 101, 23, 1}))
}
