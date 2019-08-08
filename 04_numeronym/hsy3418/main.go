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

// numeronyms returns a slice of numeronyms for its input strings.
func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	copy(result, vals)
	for index, val := range result {
		strRune := []rune(val)
		runeLength := len(strRune)
		if runeLength > 3 {
			result[index] = fmt.Sprintf("%c%d%c", strRune[0], runeLength-2, strRune[runeLength-1])
		}
	}
	return result
}
