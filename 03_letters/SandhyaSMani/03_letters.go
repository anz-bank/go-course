package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, sortLetters(letters("aba")))
}

func sortLetters(m map[rune]int) []string {
	keysArray := []int{}
	for key := range m {
		keysArray = append(keysArray, int(key))
	}
	sort.Ints(keysArray)

	var strArray = []string{}
	for _, key := range keysArray {
		strArray = append(strArray, fmt.Sprintf("%s : %d", string(key), m[rune(key)]))
	}
	return strArray

}

func letters(s string) map[rune]int {
	map1 := make(map[rune]int)
	for _, val := range s {
		map1[val]++
	}
	return map1
}
