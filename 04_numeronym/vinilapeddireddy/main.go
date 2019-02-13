package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(s ...string) []string {
	var numeronyms []string

	for _, i := range s {
		target := strings.Trim(i, "")
		len := len(target)
		if len <= 3 {
			numeronyms = append(numeronyms, target)
		} else {
			numeronyms = append(numeronyms, fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1]))
		}
	}
	return numeronyms
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
