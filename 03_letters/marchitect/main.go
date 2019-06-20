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
	// keep a list of what characters exist in the string s
	chars := ""
	for _, c := range s {
		if !strings.ContainsAny(chars, string(c)) {
			numberOfc := strings.Count(s, string(c))
			m[c] = numberOfc
			chars += string(c)
		}
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	result := make([]string, 0)
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		unquoted, _ := strconv.Unquote(strconv.QuoteRune(rune(k)))
		result = append(result, unquoted+":"+strconv.Itoa(m[rune(k)]))
	}
	return result
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
