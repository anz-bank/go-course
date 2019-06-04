package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func letters(s string) map[rune]int {
	freqMap := map[rune]int{}
	ss := []rune(s)
	for i := range ss {
		freqMap[ss[i]]++
	}
	return freqMap
}

type letterFreq struct {
	letter rune
	count  int
}

func (lf letterFreq) String() string {
	return fmt.Sprintf("%c:%d", lf.letter, lf.count)
}

func sortLetters(m map[rune]int) []string {
	freq, i := make([]letterFreq, len(m)), 0
	for l, count := range m {
		lf := letterFreq{letter: l, count: count}
		freq[i] = lf
		i++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].count > freq[j].count
	})
	sortedLetters := make([]string, len(freq))
	for i, lf := range freq {
		sortedLetters[i] = lf.String()
	}
	return sortedLetters
}

var out io.Writer = os.Stdout

func main() {
	text := "aba"
	fmt.Fprintln(out, strings.Join(sortLetters(letters(text)), "\n"))
}
