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
	letterMap := make(map[rune]int)
	for _, char := range s {
		letterMap[char]++
	}
	return letterMap

}

func sortLetters(m map[rune]int) []string {
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	sortedKey := make([]string, len(keys))
	for pos, k := range keys {
		sortedKey[pos] = fmt.Sprintf("%c:%d", k, m[rune(k)])
	}
	return sortedKey
}
