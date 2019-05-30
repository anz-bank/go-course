package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func letters(s string) map[rune]int {
	/*
		Returns a mapping of each letter to its frequency
	*/
	dict := make(map[rune]int)
	for _, c := range s {
		_, has := dict[c]
		if has {
			dict[c]++
		} else {
			dict[c] = 1
		}
	}
	return dict
}

func sortLetters(m map[rune]int) []string {
	/*
		Returns a sorted slice of strings with elements "key:val"
	*/
	// Flatten map into a string
	s := make([]string, 0)
	for k, v := range m {
		s = append(s, string(k)+":"+strconv.Itoa(v))
	}
	return sort.StringSlice(s)
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
