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
	charMap := map[rune]int{}

	for _, char := range s {
		charMap[char]++
	}

	return charMap
}

func sortLetters(m map[rune]int) []string {
	out := make([]string, len(m))
	keys := make([]int, len(m))
	index := 0

	for char := range m {
		keys[index] = int(char)
		index++
	}

	sort.Ints(keys)

	for i, key := range keys {
		out[i] = fmt.Sprintf("%c:%d", key, m[rune(key)])
	}

	return out
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
