package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func letters(s string) map[rune]int {
	var lettersMap = make(map[rune]int)
	for i, r := range s {
		i = 0
		i = strings.Count(s, string(r))
		lettersMap[r] = i
	}
	return lettersMap
}

func sortLetters(m map[rune]int) []string {
	stringKey := []string{}
	var s string
	for key, value := range m {
		s = string(key) + ":" + strconv.Itoa(value)
		stringKey = append(stringKey, s)
	}
	sort.Strings(stringKey)
	return stringKey
}
func main() {
	var m = letters("helloworld")
	for key, value := range m {
		fmt.Println("Key:", string(key), "Value:", value)
	}
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}
