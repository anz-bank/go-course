package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	return count
}

func sortLetters(m map[rune]int) []string {
	output := []string{}
	for rune, count := range m {
		output = append(output, string(rune) + ":" + strconv.Itoa(count))
	}
	sort.Strings(output)
	return output
}
