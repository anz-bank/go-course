package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	s := make([]string, len(vals))
	for i, val := range vals {
		s[i] = numeronym(val)
	}
	return s
}

func numeronym(val string) string {
	runes := []rune(strings.Trim(val, "\n\t "))
	l := len(runes)
	if l < 4 {
		return string(runes)
	}
	return fmt.Sprintf("%c%d%c", runes[0], l-2, runes[l-1])
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc", "abcd"))
}
