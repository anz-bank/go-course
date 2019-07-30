package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var mainout io.Writer = os.Stdout

func main() {

	fmt.Fprintf(mainout, strings.Join(sortLetters(letters("aba")), "\n"))

}

func letters(s string) map[rune]int {

	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {

	lettersCount := make([]string, len(m))
	i := 0
	for r := range m {
		lettersCount[i] = string(r) + ":" + strconv.Itoa(m[r])
		i++
	}
	sort.Strings(lettersCount)
	return lettersCount
}
