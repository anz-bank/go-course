package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	stats := make(map[rune]int)
	for pos, char := range s {
		stats[char] = pos
	}
	return stats
}

func sortLetters(m map[rune]int) []string {
	var keys []rune
	for k := range m {
		keys = append(keys, k)
	}
	sorted(keys)
	sortedArray := make([]string, len(keys))
	i := 0
	for _, k := range keys {
		sortedArray[i] = string(k)
		i++
	}
	return sortedArray
}

func main() {
	printRuneMap(letters("aba"))
}

func printRuneMap(m map[rune]int) {
	for key, value := range m {
		fmt.Fprintf(out, "%s:%d\n", string(key), value)
	}
}

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(runes []rune) string {
	sort.Sort(runeSlice(runes))
	return string(runes)
}
