package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

var out io.Writer = os.Stdout

func numeronyms(s ...string) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = numeronym(v)
	}

	return r
}

func numeronym(s string) string {
	c := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			c = append(c, r)
		}
	}
	l := len(c)
	if l < 4 {
		return string(c)
	}

	return fmt.Sprintf("%c%d%c", c[0], l-2, c[l-1])
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
