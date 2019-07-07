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

func letterFreq(s string) map[rune]int {
	f := map[rune]int{}
	for _, r := range s {
		f[r]++
	}
	return f
}

// Concurrent implementation of letters for speed
// Returns the letter frequency for string
func letters(s string) map[rune]int {
	stringSlice := strings.Fields(s)
	combineFreq := map[rune]int{}
	size := len(stringSlice)
	results := make(chan map[rune]int, size)

	for _, ss := range stringSlice {
		go func(sx string) {
			results <- letterFreq(sx)
		}(ss)
	}

	//Combine the letter frequency maps passed back via the results channel
	for i := 0; i < size; i++ {
		for r, freq := range <-results {
			combineFreq[r] += freq
		}
	}
	return combineFreq
}

func sortLetters(m map[rune]int) []string {
	keys := make([]string, 0, len(m))
	for key, value := range m {
		keys = append(keys, string(key)+":"+strconv.Itoa(value))
	}
	sort.Strings(keys)
	return keys
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
