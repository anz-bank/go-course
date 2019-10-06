package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	r := make([]string, len(vals))
	for k, v := range vals {
		r[k] = numeronym(v)
	}
	return r
}

func numeronym(s string) string {
	numrunes := utf8.RuneCountInString(s)
	if numrunes < 4 {
		return s
	}
	firstrune, _ := utf8.DecodeRuneInString(s)
	lastrune, _ := utf8.DecodeLastRuneInString(s)

	return string(firstrune) + strconv.Itoa(numrunes-2) + string(lastrune)
}
