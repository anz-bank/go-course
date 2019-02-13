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
	letterMap := make(map[rune]int)
	for _, r := range s {
		letterMap[r]++
	}
	return letterMap
}

func sortLetters(m map[rune]int) []string {
	//Sort keys first
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	sortedKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		sortedKeys = append(sortedKeys, fmt.Sprintf("%s:%d", string(k), m[rune(k)]))
	}
	return sortedKeys
}

func main() {
	fmt.Fprint(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
