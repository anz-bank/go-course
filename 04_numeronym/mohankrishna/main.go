package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	numeronyms := make([]string, len(vals))
	for pos, value := range vals {
		trimmedVal := []rune(strings.TrimSpace(value))
		length := len(trimmedVal)
		if length > 3 {
			numeronyms[pos] = fmt.Sprintf("%c%d%c", trimmedVal[0], length-2, trimmedVal[length-1])
		} else {
			numeronyms[pos] = string(trimmedVal)
		}
	}
	return numeronyms
}
