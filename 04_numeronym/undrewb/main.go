package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	strs := make([]string, len(vals))
	for i, s := range vals {
		strs[i] = numeronym(s)
	}
	return strs
}

func numeronym(s string) string {
	if len([]rune(s)) > 3 {
		r := []rune(s)
		s = fmt.Sprintf("%c%d%c", r[0], len(r)-2, r[len(r)-1])
	}
	return s
}
