package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	x := make(map[rune]int)
	for _, char := range s {
		x[char]++
	}
	return x
}

func sortLetters(m map[rune]int) []string {
	sorted := []string{}
	for i, val := range m {
		sorted = append(sorted, fmt.Sprintf("%c:%v", i, val))
	}
	sort.Strings(sorted)
	return sorted
}
