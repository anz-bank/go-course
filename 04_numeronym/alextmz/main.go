package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	r := make([]string, len(vals))
	for k, v := range vals {
		r[k] = numeronym(v)
	}
	return r
}

func numeronym(s string) string {
	if len(s) < 4 {
		return (s)
	}
	return s[:1] + strconv.Itoa(len(s)-2) + s[len(s)-1:]
}
