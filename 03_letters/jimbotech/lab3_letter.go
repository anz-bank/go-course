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
	letterCounter := map[rune]int{}

	for _, c := range s {
		letterCounter[c]++
	}
	return letterCounter
}

func sortLetters(m map[rune]int) []string {
	kvs := make([]string, 0, len(m))
	for key, val := range m {
		kvs = append(kvs, fmt.Sprintf("%c:%d", key, val))
	}
	sort.Strings(kvs)
	return kvs
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
