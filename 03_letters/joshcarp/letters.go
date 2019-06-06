package main

import (
	"fmt"
	"sort"
	"strings"
)

// Code from https://www.socketloop.com/tutorials/golang-sort-and-reverse-sort-a-slice-of-runes
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, i := range s {
		freq[i]++
	}
	return freq
}

func sortLetters(m map[rune]int) []string {
	var ans []string
	var keys RuneSlice

	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Sort(keys)
	for _, key := range keys {
		ans = append(ans, fmt.Sprintf("%c:%d", key, m[rune(key)]))
	}
	return ans
}
