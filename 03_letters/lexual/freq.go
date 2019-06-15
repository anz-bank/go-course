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
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = string(k)
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(m))
	for i, k := range keys {
		sorted[i] = fmt.Sprintf("%s:%d", k, m[rune(k[0])])
	}
	return sorted
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
