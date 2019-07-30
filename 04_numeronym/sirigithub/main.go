package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	for i, val := range vals {
		r := []rune(val)
		runeLength := len(r)
		if runeLength > 3 {
			val = fmt.Sprintf("%c%d%c", r[0], runeLength-2, r[runeLength-1])
		}
		result[i] = val
	}
	return result
}
func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
