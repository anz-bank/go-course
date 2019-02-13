package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout

func main() {
	// fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
	// sortLetters(letters("aba"))
	fmt.Println(sortLetters(letters("aba")))
}

func letters(s string) map[rune]int {
	letterMap := make(map[rune]int)
	for _, v := range s {
		letterMap[v]++
	}
	return letterMap
}

func sortLetters(m map[rune]int) []rune {
	// var sortedMap []string

	sortedMap := make([]rune, 0, len(m))

	for v := range m {
		sortedMap = append(sortedMap, v)
	}
	sort.Ints(sortedMap)
	for _, name := range sortedMap {
		//...
	}

	return sortedMap
}
