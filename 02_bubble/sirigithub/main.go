package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		for j :=1 ; j < n-i; j++ {

			if s[j-1] > s[j] {
				s[j-1], s[j]= s[j], s[j-1]
			}
		}
    fmt.Println("   ",s)
	}
	return s
}

func insertion(s []int) []int {
  i:= 1

  for i < len(s) {
    j:= i

    for j > 0 && s[j-1] > s[j] {
      s[j-1] , s[j] = s[j], s[j-1]
      j--
    }
    i++
  }
 return s

}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
  fmt.Fprintln(out, insertion([]int{8, 2, 1, 6}))
}
