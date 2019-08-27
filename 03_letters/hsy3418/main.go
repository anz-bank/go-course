package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

//letters is a function returns a mapping of each letter to its frequency.
func letters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

//sorLetters is a function returns a sorted slice of strings with elements
func sortLetters(m map[rune]int) []string {
	sortSlices := make([]string, 0, len(m))
	for key, val := range m {
		sortSlices = append(sortSlices, fmt.Sprintf("%c:%d", key, val))
	}
	sort.Strings(sortSlices)
	return sortSlices
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
