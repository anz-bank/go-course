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
		out[i] = numeronym(str)
	}
	return out
}

func numeronym(input string) string {
	trimmedInput := strings.TrimSpace(input)

	//  Unicode string could not be accessed using index
	//  Therefore, string is converted into an array of runes
	runes := []rune(trimmedInput)
	length := len(runes)

	if length <= 3 {
		return trimmedInput
	}

	return fmt.Sprintf("%c%d%c", runes[0], length-2, runes[length-1])
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
