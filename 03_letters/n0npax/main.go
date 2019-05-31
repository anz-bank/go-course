package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func letters(s string) (freqMap map[rune]int) {
	freqMap = map[rune]int{}
	ss := []rune(s)
	for i := range ss {
		freqMap[ss[i]]++
	}
	return
}

type letterFreq struct {
	letter    rune
	occurency int
}

func (lf letterFreq) String() string {
	return fmt.Sprintf("%c:%d", lf.letter, lf.occurency)
}

func sortLetters(m map[rune]int) (sortedLetters []string) {

	freq, i := make([]letterFreq, len(m)), 0
	for l, occurency := range m {
		lf := letterFreq{letter: l, occurency: occurency}
		freq[i] = lf
		i++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].occurency > freq[j].occurency
	})
	sortedLetters = make([]string, len(freq))
	for i, lf := range freq {
		sortedLetters[i] = lf.String()
	}
	return
}

var out io.Writer = os.Stdout

func main() {
	text := "aba"
	fmt.Fprintln(out, strings.Join(sortLetters(letters(text)), "\n"))
}
