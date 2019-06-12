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
	letterMap := make(map[rune]int)

	for _, value := range s {
		letterMap[value]++
	}
	return letterMap
}

func sortLetters(m map[rune]int) []string {
	frequencySlice := make([]string, 0, len(m))

	for k, v := range m {
		keyInStringFormat := string(k)
		mapIntToString := strconv.Itoa(v)
		frequencyOutput := keyInStringFormat + ":" + mapIntToString
		frequencySlice = append(frequencySlice, frequencyOutput)
	}
	sort.Strings(frequencySlice)
	return frequencySlice
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
