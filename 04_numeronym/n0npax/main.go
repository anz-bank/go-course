package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func shorten(word string) string {
	if len(word) < 4 {
		return word
	}
	num := len(word) - 2
	return fmt.Sprintf("%c%d%c", word[0], num, word[len(word)-1])
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
