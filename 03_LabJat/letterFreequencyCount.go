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

func letters(input string) map[rune]int {
	m := map[rune]int{}
	for _, c := range input {
		m[c]++
	}
	return m
}
func sortedLetters(m map[rune]int) []string {
	var keys []string
	for k, v := range m {
		keys = append(keys, string(k)+":"+strconv.Itoa(v))
	}
	sort.Strings(keys)
	return keys
}
func main() {
	fmt.Fprintln(out, strings.Join(sortedLetters(letters("This is a String")), "\n"))
}
