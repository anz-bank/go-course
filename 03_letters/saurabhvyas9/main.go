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
	frequencyLetters := make(map[rune]int)
	for _, str := range s {
		frequencyLetters[str]++
	}
	return frequencyLetters
}

func sortLetters(m map[rune]int) []string {
	sortedLetters := make([]string, len(m))
	tempKey := make([]int, len(m))
	i := 0
	for key := range m {
		tempKey[i] = int(key)
		i++
	}
	sort.Ints(tempKey)

	for index, runeKey := range tempKey {
		sortedLetters[index] = fmt.Sprintf("%c:%d", runeKey, m[rune(runeKey)])
	}
	return sortedLetters
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
