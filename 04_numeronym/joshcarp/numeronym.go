package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	result := make([]string, 0, len(vals))
	for _, word := range vals {
		word = removeWhitespaces(word)
		if length := len(word); length > 3 && isWord(word) {
			word = fmt.Sprintf("%s%d%s", string(word[0]), length-2, string(word[length-1]))
		}
		result = append(result, word)
	}
	return result
}

func removeWhitespaces(word string) string {
	var newWord string
	for _, letter := range word {
		if letter != ' ' {
			newWord += string(letter)
		}
	}
	return newWord
}
func isWord(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			return false
		}
	}
	return true
}
