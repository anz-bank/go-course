package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	cnt := len(s)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if s[i] > s[j] {
				tmp := s[i]
				s[i], s[j] = s[j], tmp
				j--
			}
		}
	}
	return s
}

func main() {
	s := bubble([]int{3, 2, 1, 5})
	v := strings.Join(strings.Fields(fmt.Sprint(s)), " ")
	fmt.Fprintln(out, v)
}
