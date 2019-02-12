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
	fmt.Fprint(out, strings.Join(sortLetters(countLetters("aba")), "\n"), "\n")
}

func countLetters(s string) map[rune]int {
	elementMap := map[rune]int{}
	for _, c := range s {
		elementMap[c]++
	}
	return elementMap
}

func sortLetters(m map[rune]int) []string {
	keys := []int{}
	sortedOutput := []string{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		sortedOutput = append(sortedOutput, fmt.Sprintf("%c : %d", k, m[rune(k)]))
	}
	return sortedOutput
}
