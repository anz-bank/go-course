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
	m := map[rune]int{}
	for _, val := range s {
		value := m[val]
		m[val] = value + 1
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
