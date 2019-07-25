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
	count := map[rune]int{}
	for _, char := range s {
		count[char]++
	}
	return count
}

func sortLetters(m map[rune]int) []string {
	result := make([]string, len(m))
	keys := []int{}

	// loop to get keys out of m. Ranging through m to extract keys to fill up keys slice
	for char := range m {
		keys = append(keys, int(char))
	}
	sort.Ints(keys)
	// building the result []string
	for i, key := range keys {
		result[i] = fmt.Sprintf("%c:%d", key, m[rune(key)])
		// note that Sprintf returns u a formatted string it doesn't print anything
	}
	return result
}

func main() {
	// the strings.Join is to concatenate the new space to every char in s
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
