package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func numeronym(val string) string {
	l := utf8.RuneCountInString(val)
	if l < 4 {
		return val
	}
	firstRune, _ := utf8.DecodeRuneInString(val)
	lastRune, _ := utf8.DecodeLastRuneInString(val)
	return fmt.Sprintf("%v%v%v", string(firstRune), l-2, string(lastRune))
}

func numeronyms(vals ...string) []string {
	numero := []string{}
	for _, v := range vals {
		numero = append(numero, numeronym(v))
	}
	return numero
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
