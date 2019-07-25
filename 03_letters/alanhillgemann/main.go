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
	letterTotalMap := map[rune]int{}
	for _, r := range s {
		letterTotalMap[r]++
	}
	return letterTotalMap
}

func sortLetters(m map[rune]int) []string {
	letterTotalStrings := []string{}
	for k, v := range m {
		letterTotalStrings = append(letterTotalStrings, fmt.Sprintf("%s:%d", string(k), v))
	}
	sort.Strings(letterTotalStrings)
	return letterTotalStrings
}
