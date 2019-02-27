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
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}

	return m
}

func sortLetters(m map[rune]int) []string {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	s := make([]string, 0, len(m))
	for _, k := range keys {
		s = append(s, fmt.Sprintf("%c:%d", k, m[rune(k)]))
	}

	return s
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
