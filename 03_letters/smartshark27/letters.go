package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
	"strconv"
)

func letters(s string) map[rune]int {
	m := make(map[rune]int)
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		m[runeValue]++
		w = width
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	sl := make([]string, len(m))
	ind := 0
	for letter, count := range m {
		sl[ind] = string(letter) + ":" + strconv.Itoa(count)
		ind++
	}

	sort.Strings(sl)
	return sl
}

func main() {
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}