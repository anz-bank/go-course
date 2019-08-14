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

// Returns the letter frequency for a string
func letters(s string) map[rune]int {
	str := strings.Fields(s)
	size := len(str)
	results := make(chan map[rune]int, size)

	for _, ss := range str {
		go func(sx string) {
			results <- letterFreq(sx)
		}(ss)
	}

	sum := map[rune]int{}
	// Combine the letter frequency maps passed back via the results channel
	for i := 0; i < size; i++ {
		result := <-results
		for r, freq := range result {
			sum[r] += freq
		}
	}
	return sum
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
