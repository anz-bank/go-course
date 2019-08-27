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
	for _, v := range s {
		m[v]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	s := []string{}
	for k, v := range m {
		s = append(s, fmt.Sprintf("%c:%d", k, v))
	}
	sort.Strings(s)
	return s
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
