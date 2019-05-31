package main

import (
	"fmt"
	"io"
	"os"
)

var outbub io.Writer = os.Stdout

func bubble(s []int) []int {
	temp := make([]int, len(s))
	copy(temp, s)
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if temp[j] > temp[j+1] {
				// swap arr[j+1] and arr[i]
				tmp := temp[j]
				temp[j] = temp[j+1]
				temp[j+1] = tmp
			}
		}
	}
	return temp
}
func insertion(s []int) []int {
	temp := make([]int, len(s))
	copy(temp, s)
	n := len(s)

	for i := 0; i < n-1; i++ {
		key := temp[i]
		j := i - 1

		for ; j >= 0 && temp[j] > key; j-- {
			temp[j+1] = temp[j]
		}
		temp[j+1] = key
	}
	return temp
}
func main() {

	fmt.Fprint(outbub, bubble([]int{3, 2, 1, 5}))
	fmt.Fprint(outbub, insertion([]int{3, 2, 1, 5}))
}
