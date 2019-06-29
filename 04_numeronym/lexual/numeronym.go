package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronym(s string) string {
	if len(s) <= 3 {
		return s
	}
	return s[:1] + strconv.Itoa(len(s)-2) + s[len(s)-1:]
}

func numeronyms(vals ...string) []string {
	ns := make([]string, len(vals))
	for i := range vals {
		ns[i] = numeronym(vals[i])
	}
	return ns
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
