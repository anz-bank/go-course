package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, getNumeronyms("accessibility", "Kubernetes", "abc"))
}

func getNumeronyms(s ...string) []string {
	numeronyms := make([]string, len(s))
	for index, i := range s {
		target := []rune(strings.TrimSpace(i))
		len := len(target)
		if len <= 3 {
			numeronyms[index] = string(target)
		} else {
			numeronyms[index] = fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1])
		}
	}
	return numeronyms
}
