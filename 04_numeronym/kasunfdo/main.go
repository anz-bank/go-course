package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(s ...string) []string {
	out := make([]string, len(s))
	for i, str := range s {
		out[i] += getNumeronym(str)
	}
	return out
}

func getNumeronym(str string) string {
	trimmedString := strings.TrimSpace(str)

	//  Unicode string could not be accessed using index
	//  Therefore, string is converted into an array of runes
	runeArr := []rune(trimmedString)
	length := len(runeArr)

	if length <= 3 {
		return trimmedString
	}

	return fmt.Sprintf("%c%d%c", runeArr[0], length-2, runeArr[length-1])
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
