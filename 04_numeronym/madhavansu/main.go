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
	var strcopy []string
	for _, val := range vals {
		val := []rune(strings.TrimSpace(val))
		length := len(val)
		if length > 3 {
			strcopy = append(strcopy, fmt.Sprint(string(val[0:1]), length-2, string(val[length-1:])))
		} else {
			strcopy = append(strcopy, string(val))
		}
	}
	return strcopy
}
