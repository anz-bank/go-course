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
			s, _ := utf8.DecodeRuneInString(val)     // starting character
			e, _ := utf8.DecodeLastRuneInString(val) // ending character
			val = fmt.Sprintf("%c%d%c", s, l-2, e)
		}
		result = append(result, val)
	}
	return result
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
