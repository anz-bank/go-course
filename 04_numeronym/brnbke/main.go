package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var result = make([]string, 0, len(vals))
	for _, val := range vals {
		l := utf8.RuneCountInString(val)
		if l > 3 {
			s := []rune(val)[0]   // starting character
			e := []rune(val)[l-1] // ending character
			result = append(result, fmt.Sprintf("%c%d%c", s, l-2, e))
		} else {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
