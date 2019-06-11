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
		if letterFrequency, found := letterMap[value]; found {
			letterMap[value] = letterFrequency + 1
		} else {
			letterMap[value] = 1
		}
	}
	return letterMap
}

func sortLetters(m map[rune]int) []string {
	frequencySlice := []string{}

	for k, v := range m {
		convertedKeyInStringFormat := string(k)
		mapIntToString := strconv.Itoa(v)
		frequencyOutput := convertedKeyInStringFormat + ":" + mapIntToString
		frequencySlice = append(frequencySlice, frequencyOutput)
	}
	sort.Strings(frequencySlice)
	return frequencySlice
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
