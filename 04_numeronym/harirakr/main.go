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

func numeronym(w string) string {
	rw := []rune(w)
	if n := len(rw); n > 3 {
		w = fmt.Sprintf("%c%d%c", rw[0], n-2, rw[n-1])
	}
	return w
}

func numeronyms(vals ...string) []string {
	s := []string{}
	for _, v := range vals {
		s = append(s, numeronym(v))
	}
	return s
}
