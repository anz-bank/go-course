package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	out := map[rune]int{}
	for _, runeValue := range s {
		out[runeValue]++
	}
	return out
}

func sortLetters(m map[rune]int) []string {
	out := make([]string, 0, len(m))
	for r, c := range m {
		str := fmt.Sprintf("%c:%v", r, c)
		out = append(out, str)
	}
	sort.Strings(out)
	return out
}
