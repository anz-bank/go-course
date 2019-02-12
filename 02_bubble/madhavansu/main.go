package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	arrlen := len(arr)
	for i := 0; i < arrlen-1; i++ {
		for j := 0; j < arrlen-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func insertion(arrcopy []int) []int {
	arr := make([]int, len(arrcopy))
	copy(arr, arrcopy)
	arrlen := len(arr)
	for i := 1; i < arrlen; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
