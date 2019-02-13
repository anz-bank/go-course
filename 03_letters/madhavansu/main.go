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
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	s := make([]string, len(keys))
	for i, v := range keys {
		s[i] = fmt.Sprintf("%c:%d", v, m[rune(v)])
	}
	return s
}
