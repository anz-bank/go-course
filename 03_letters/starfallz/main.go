package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, currentLetter := range s {
		letterCount[currentLetter]++
	}

	return letterCount
}

func sortLetters(m map[rune]int) []string {
	allKeys := make([]rune, 0, len(m))
	for k := range m {
		allKeys = append(allKeys, k)
	}

	sortRune(allKeys)

	result := make([]string, len(allKeys))

	for i, currentKey := range allKeys {
		result[i] = strings.Join([]string{string(currentKey), strconv.Itoa(m[currentKey])}, ":")
	}

	return result
}

func sortRune(r []rune) {
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
}
