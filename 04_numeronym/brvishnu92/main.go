package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	x := make([]string, len(vals))
	for i, val := range vals {
		if len(val) > 3 {
			x[i] = fmt.Sprintf("%c%d%c", val[0], len(val)-2, val[len(val)-1])
		} else {
			x[i] = val
		}
	}
	return x
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
