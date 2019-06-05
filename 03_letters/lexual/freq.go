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
	var keys []rune
	for k := range m {
		keys = append(keys, k)
	}
	// sort keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	var sorted []string
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s:%d", string(k), m[k]))
	}
	return sorted
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
