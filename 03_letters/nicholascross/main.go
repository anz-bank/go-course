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
	keys := make([]string, 0, len(m))
	for k, v := range m {
		keys = append(keys, fmt.Sprintf("%c:%d", k, v))
	}
	sort.Strings(keys)
	return keys
}

func letters(s string) map[rune]int {
	count := map[rune]int{}

	for _, char := range s {
		count[char]++
	}
	return count
}
