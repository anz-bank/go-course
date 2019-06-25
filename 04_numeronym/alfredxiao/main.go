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
		result[i] = numeronym(val)
	}
	return result
}

func numeronym(val string) string {
	runes := []rune(val)
	length := len(runes)
	if length > 3 {
		return fmt.Sprintf("%c%d%c", runes[0], length-2, runes[length-1])
	}
	return val
}

func main() {
	fmt.Fprint(out, numeronyms("abcd"))
}
