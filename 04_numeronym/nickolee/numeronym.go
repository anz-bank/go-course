package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	for i, val := range vals {
		result[i] = numeronym(val) // build results slice
	}
	return result
}

func numeronym(s string) string {
	// convert string to slice of runes (Unicode int32) to avoid some kind of "protocol error"
	runes := []rune(s)
	if len(runes) <= 3 {
		return string(runes)
	}

	firstChar := string(runes[0]) // casting a byte (one character) into a string
	lastChar := string(runes[len(runes)-1])
	middleChars := strconv.Itoa(len(runes) - 2)
	return firstChar + middleChars + lastChar
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
