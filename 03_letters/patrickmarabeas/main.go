package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

// Letters creates a frequency map of runes within the given string.
func letters(s string) map[rune]int {
	result := make(map[rune]int)
	// Increment key's value by one for each instance
	// Starts with the zero value of the ValueType (int)
	for _, v := range s {
		result[v]++
	}

	return result
}

// SortLetters creates a slice of key value pairs in the format of "{unicode character}:{frequency}"
// sorted by ascending key.
func sortLetters(m map[rune]int) []string {
	length := len(m)
	result := make([]string, length)
	// Create a slice of runes from map keys
	runes := make([]rune, length)
	i := 0
	for k := range m {
		runes[i] = k
		i++
	}
	// Sort the slice of runes
	// The less function that sort implements takes int32 which aligns with rune
	// https://golang.org/src/sort/slice.go?s=451:506#L7
	// https://golang.org/src/runtime/sema.go?h=less#L468
	// https://golang.org/src/builtin/builtin.go?h=rune#L92
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	//
	for i, v := range runes {
		// Format as "{unicode character}:{decimal}"
		result[i] = fmt.Sprintf("%c:%d", v, m[v])
	}

	return result
}

func main() {
	fmt.Fprint(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
