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
	length := len(val)
	if length > 3 {
		return fmt.Sprintf("%c%d%c", val[0], length-2, val[length-1])
	}
	return val
}

func main() {
	fmt.Fprint(out, numeronyms("abcd"))
}
