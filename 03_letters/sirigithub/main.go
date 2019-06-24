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

// letters returns a map of rune literals and its frequency in the given input string s
func letters(s string) map[rune]int {
	charecterFreq := make(map[rune]int)
	// increment the counter if the charecter exists
	for _, character := range s {
		charecterFreq[character]++
	}
	return charecterFreq
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
