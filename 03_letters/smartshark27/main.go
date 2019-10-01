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

func letters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
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
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
