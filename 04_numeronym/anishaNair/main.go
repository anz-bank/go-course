package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(s ...string) []string {
	numeronyms := make([]string, len(s))
	for index, i := range s {
		runes := []rune(strings.Trim(i, "\n\t "))
		len := len(runes)
		if len <= 3 {
			numeronyms[index] = string(runes)
		} else {
			numeronyms[index] = fmt.Sprintf("%c%d%c", runes[0], len-2, runes[len-1])
		}
	}
	return numeronyms
}
