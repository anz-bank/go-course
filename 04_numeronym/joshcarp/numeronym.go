package main

import "fmt"

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	result := []string{}
	for _, word := range vals {
		newWord := word
		if length := len(word); length > 3 {
			newWord = fmt.Sprintf("%s%d%s", string(word[0]), length-2, string(word[length-1]))
		}
		result = append(result, newWord)
	}
	return result
}
