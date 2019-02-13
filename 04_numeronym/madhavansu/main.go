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
	strcopy := make([]string, len(vals))
	for k, val := range vals {
		val := []rune(strings.TrimSpace(val))
		length := len(val)
		if length > 3 {
			strcopy[k] = fmt.Sprint(string(val[0]), length-2, string(val[length-1]))
		} else {
			strcopy[k] = string(val)
		}
	}
	return strcopy
}
