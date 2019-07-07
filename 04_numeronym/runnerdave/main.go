package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	n := make([]string, len(vals))
	for i, value := range vals {
		length := utf8.RuneCountInString(value)
		if length > 3 {
			n[i] = fmt.Sprintf("%v%d%v", string(value[0]), length-2, string(value[length-1]))
		} else {
			n[i] = value
		}
	}
	return n
}
