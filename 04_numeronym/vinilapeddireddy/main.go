package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(s ...string) []string {
	numeronyms := make([]string, len(s))
	for i, j := range s {
		target := strings.Trim(j, " ")
		len := len(target)
		if len <= 3 {
			numeronyms[i] = target
		} else {
			numeronyms[i] = fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1])
		}
	}
	return numeronyms
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
