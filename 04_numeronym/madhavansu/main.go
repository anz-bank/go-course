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
	var strcopy []string
	for _, val := range vals {
		length := len(val)
		if length > 3 {
			first, last := val[0:1], val[length-1:]
			strcopy = append(strcopy, fmt.Sprint(first, length-2, last))
		} else {
			strcopy = append(strcopy, val)
		}

	}
	return strcopy
}
