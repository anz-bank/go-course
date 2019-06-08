package main

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

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
	ans := make(StringSlice, 0, len(m))
	for key, val := range m {
		ans = append(ans, fmt.Sprintf("%c:%d", key, val))
	}
	sort.Sort(ans)
	return ans
}
