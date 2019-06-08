package main

import "fmt"

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	result := make([]string, 0, len(vals))
	for _, word := range vals {
		if length := len(word); length > 3 {
			word = fmt.Sprintf("%s%d%s", string(word[0]), length-2, string(word[length-1]))
		}
		result = append(result, word)
	}
	return result
}
