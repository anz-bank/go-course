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

// letters returns a map of rune literals and its frequency in the given string s
func letters(s string) map[rune]int {

	charecterFreq := make(map[rune]int)
	// increment the counter if the charecter exists
	for _, character := range s {
		charecterFreq[character]++
	}
	fmt.Println(charecterFreq)
	return charecterFreq
}

func sortLetters(m map[rune]int) []string {
	sortedStrings := make([]string, len(m))

	for key, value := range m {
		str := string(key) + ":" + strconv.Itoa(value)
		sortedStrings = append(sortedStrings, str)
	}
	sort.Strings(sortedStrings)

	return sortedStrings
}

func main() {
	fmt.Println(out, strings.Join(sortLetters(letters("dasdsadsd")), "\n"))
}
