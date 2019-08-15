package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronym(r []rune) string {
	return fmt.Sprintf("%c%d%c", r[0], len(r)-2, r[len(r)-1])
}

func numeronyms(vals ...string) []string {
	var s = make([]string, len(vals))
	copy(s, vals)

	for i, str := range s {
		r := []rune(str)
		if len(r) > 3 {
			s[i] = numeronym(r)
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
