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
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	s := make([]string, len(m))
	i := 0
	for k, v := range m {
		s[i] = string(k) + ":" + strconv.Itoa(v)
		i++
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}
