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
	fmt.Fprintln(out, strings.Join(sortLetters(lettersFreq("aba")), "\n"))
}

func lettersFreq(s string) map[rune]int {
	lettersMap := make(map[rune]int)
	for _, letter := range s {
		lettersMap[letter]++
	}
	return lettersMap
}

func sortLetters(m map[rune]int) []string {
	sorted := []string{}
	for index, val := range m {
		sorted = append(sorted, fmt.Sprintf("%c:%v", index, val))
	}
	sort.Strings(sorted)
	return sorted
}
