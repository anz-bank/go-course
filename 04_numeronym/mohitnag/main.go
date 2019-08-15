package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	ns := make([]string, len(vals))
	for i, val := range vals {
		ns[i] = numeronym(val)
	}
	return ns
}

func numeronym(v string) string {
	rs := []rune(v)
	length := len(rs)
	if length <= 3 {
		return v
	}
	return fmt.Sprintf("%c%d%c", rs[0], length-2, rs[length-1])
}
