package merge

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//merge sort: O(n log n) sorting algorithm
func merge(s []int) []int {
	sBox := [][]int{}
	for a := range s {
		sBox = append(sBox, []int{s[a]})
	}

	for len(sBox) > 1 {
		var tBox [][]int
		for i := 0; i < len(sBox); i += 2 {
			//if the adjacent box exists then compare them
			if i+1 != len(sBox) {
				tSlice := []int{}
				//need to do for loop here to check if i and i+1 lengths are greater than 0
				for len(sBox[i]) > 0 && len(sBox[i+1]) > 0 {
					a, b := sBox[i][0], sBox[i+1][0]
					if a < b {
						tSlice = append(tSlice, a)
						sBox[i] = sBox[i][1:]
					} else {
						tSlice = append(tSlice, b)
						sBox[i+1] = sBox[i+1][1:]
					}
				}
				//need to do a check here for which array has less values
				if len(sBox[i]) > 0 {
					tSlice = append(tSlice, sBox[i]...)
				} else if len(sBox[i+1]) > 0 {
					tSlice = append(tSlice, sBox[i+1]...)
				}
				tBox = append(tBox, tSlice)
			} else {
				tBox = append(tBox, sBox[i])
			}
		}
		//set sBox to value of tBox
		sBox = tBox
	}
	return sBox[0]
}

func main() {
	fmt.Fprintln(out, merge([]int{3, 2, 1, 5}))
}
