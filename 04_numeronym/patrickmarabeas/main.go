package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// GenerateNumeronym creates a numeronym from the given string.
func generateNumeronym(s string) string {
	// Convert to runes to handle multi byte characters
	runes := []rune(s)
	length := len(runes)
	if length < 4 {
		return string(runes)
	}
	return fmt.Sprintf("%c%d%c", runes[0], length-2, runes[length-1])
}

// Numeronyms takes an input of n strings and returns a slice containing each string's numeronym.
func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	for i, numeronym := range vals {
		result[i] = generateNumeronym(numeronym)
	}
	return result
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
