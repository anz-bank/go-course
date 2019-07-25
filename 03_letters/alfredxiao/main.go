package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

func sortLetters(m map[rune]int) []string {
	rcs := []string{}
	for r, c := range m {
		rcs = append(rcs, fmt.Sprintf("%c:%d", r, c))
	}
	sort.Strings(rcs)
	return rcs
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
