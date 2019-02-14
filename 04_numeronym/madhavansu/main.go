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
	s := make([]string, len(vals))
	for k, val := range vals {
		val := []rune(strings.Trim(val, "\n\t\\ "))
		length := len(val)
		if length > 3 {
			s[k] = fmt.Sprintf("%c%d%c", val[0], length-2, val[length-1])
		} else {
			s[k] = string(val)
		}
	}
	return s
}
