package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	numeronyms := make([]string, len(vals))

	for pos, val := range vals {
		var y = []rune(val)
		if len(y) < 4 {
			numeronyms[pos] = val
		} else {
			numeronyms[pos] = fmt.Sprintf("%c%d%c", y[0], len(y)-2, y[len(y)-1])
		}
	}
	return numeronyms
}
