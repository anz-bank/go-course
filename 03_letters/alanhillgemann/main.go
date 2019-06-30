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
	runeMap := make(map[rune]int)
	runeSlice := []rune(s)
	for i := 0; i <= len(runeSlice)-1; i++ {
		runeMap[runeSlice[i]]++
	}
	return runeMap
}

func sortLetters(m map[rune]int) []string {
	stringSlice := make([]string, 0)
	mapKeys := make([]string, 0)
	mapValues := make([]int, 0)
	for k, v := range m {
		mapKeys = append(mapKeys, string(k))
		mapValues = append(mapValues, v)
	}
	for i := 0; i <= len(m)-1; i++ {
		stringSlice = append(stringSlice, fmt.Sprintf("%s:%d", mapKeys[i], mapValues[i]))
	}
	sort.Strings(stringSlice)
	return stringSlice
}
