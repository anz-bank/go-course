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
	r := make([]string, len(vals))
	for i, v := range vals {
		var val = []rune(strings.Trim(v, "\n "))
		if len(val) < 4 {
			r[i] = v
		} else {
			r[i] = fmt.Sprintf("%c%d%c", val[0], len(val)-2, val[len(val)-1])
		}
	}
	return r
}
