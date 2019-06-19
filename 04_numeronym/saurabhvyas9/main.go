package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	numeronym := make([]string, len(vals))
	for i, str := range vals {
		length := len(str)
		if length > 3 {
			numeronym[i] = fmt.Sprintf("%v%d%v", string(str[0]), length-2, string(str[length-1]))
		} else {
			numeronym[i] = str
		}
	}
	return numeronym
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
