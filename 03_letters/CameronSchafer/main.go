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
	var m = make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	return m
}

func sortLetters(m map[rune]int) []string {
	var s = []string{}

	for r, c := range m {
		s = append(s, fmt.Sprintf("%c:%v", r, c))
	}
	sort.Strings(s)

	return s
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
