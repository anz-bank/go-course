package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	numeronym := make([]string, len(vals))
	for i, str := range vals {
		strRune := []rune(str)
		runeLength := len(strRune)
		if runeLength > 3 {
			numeronym[i] = fmt.Sprintf("%c%d%c", strRune[0], runeLength-2, strRune[runeLength-1])
		} else {
			numeronym[i] = str
		}
	}
	return numeronym
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
