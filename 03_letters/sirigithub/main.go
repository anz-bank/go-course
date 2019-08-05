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
	runeFreq := make(map[rune]int)
	for _, r := range s {
		runeFreq[r]++
	}
	return runeFreq
}

// sortLetters returns a sorted slice of strings with elements {key}:{val} from the input map m
func sortLetters(m map[rune]int) []string {
	sortedStrings := make([]string, 0, len(m))
	for key, value := range m {
		frequency := string(key) + ":" + strconv.Itoa(value)
		sortedStrings = append(sortedStrings, frequency)
	}
	sort.Strings(sortedStrings)
	return sortedStrings
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
