package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	numeronyms := []string{}
	for _, v := range vals {
		n := utf8.RuneCountInString(v)
		if n > 3 {
			first, _ := utf8.DecodeRuneInString(v)
			last, _ := utf8.DecodeLastRuneInString(v)
			result := fmt.Sprintf("%v%v%v", string(first), n-2, string(last))
			numeronyms = append(numeronyms, result)
		} else {
			numeronyms = append(numeronyms, v)
		}
	}
	return numeronyms
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
