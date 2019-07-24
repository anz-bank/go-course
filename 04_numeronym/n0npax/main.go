package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func shorten(word string) string {
	wordLen := len(word)
	if wordLen < 4 {
		return word
	}
	num := wordLen - 2
	return fmt.Sprintf("%c%d%c", word[0], num, word[wordLen-1])
}

func numeronyms(vals ...string) (shortened []string) {
	shortened = make([]string, len(vals))
	for i, val := range vals {
		shortened[i] = shorten(val)
	}
	return
}

func main() {
	fmt.Fprintln(out, numeronyms("Kubernetes", "anz", "internationalization"))
}
