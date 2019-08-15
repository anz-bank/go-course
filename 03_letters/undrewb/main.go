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

func sortLetters(m map[rune]int) []string {
	strs := make([]string, 0, len(m))
	for r, c := range m {
		strs = append(strs, fmt.Sprintf("%c:%d", r, c))
	}
	sort.Strings(strs)
	return strs
}

func letters(s string) map[rune]int {
	rc := make(map[rune]int)
	for _, r := range s {
		rc[r]++
	}
	return rc
}
