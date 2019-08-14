package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	for i, val := range vals {
		c := utf8.RuneCountInString(val)
		if c > 3 {
			f, _ := utf8.DecodeRuneInString(val)
			l, _ := utf8.DecodeLastRuneInString(val)
			val = fmt.Sprintf("%c%d%c", f, c-2, l)
		}
		result[i] = val
	}
	return result
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
