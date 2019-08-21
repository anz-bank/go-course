package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	strArr := make([]string, len(vals))
	for i, str := range vals {
		strArr[i] = numeronym([]rune(str))
	}
	return strArr
}

func numeronym(word []rune) string {
	if len(word) <= 3 {
		return string(word)
	}
	count := len(word) - 2
	return fmt.Sprintf("%c%d%c", word[0], count, word[len(word)-1])
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
