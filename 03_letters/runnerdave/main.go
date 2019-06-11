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
	stats := make(map[rune]int)
	for _, char := range s {
		stats[char]++
	}
	return stats
}

func sortLetters(m map[rune]int) []string {
	keys := make([]rune, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sorted(keys)
	sortedArray := make([]string, len(keys))
	for i, k := range keys {
		sortedArray[i] = fmt.Sprintf("%s:%s", string(k), strconv.Itoa(m[k]))
	}
	return sortedArray
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(runes []rune) string {
	sort.Sort(runeSlice(runes))
	return string(runes)
}
