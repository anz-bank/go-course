package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func sortLetters(m map[rune]int) []string {
	letterFreq := []string{}
	for k, v := range m {
		s := string(k) + ":" + strconv.Itoa(v)
		letterFreq = append(letterFreq, s)
	}
	sort.Strings(letterFreq)
	return letterFreq
}

func main() {
	fmt.Fprintln(out, sortLetters(letters("This is a string for mapping!")))
}
