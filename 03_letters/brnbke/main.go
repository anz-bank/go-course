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
	result := make(map[rune]int)

	for _, r := range s {
		result[r]++
	}

	return result
}

func sortLetters(m map[rune]int) []string {
	var result = make([]string, 0, len(m))
	for k, v := range m {
		result = append(result, fmt.Sprintf("%c:%d", k, v))
	}
	sort.Strings(result)
	return result
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
