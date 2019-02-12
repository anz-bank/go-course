package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	letterMap := make(map[rune]int)

	for _, r := range s {
		if _, ok := letterMap[r]; ok {
			letterMap[r] = letterMap[r] + 1
		} else {
			letterMap[r] = 1
		}
	}
	return letterMap
}

func print(m map[rune]int) {
	//Sort keys first
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	//Print
	for _, k := range keys {
		fmt.Fprint(out, fmt.Sprintf("%s:%d\n", string(k), m[rune(k)]))

	}
}

func main() {
	result := letters("aba")
	print(result)
}
