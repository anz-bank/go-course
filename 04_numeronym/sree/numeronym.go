package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	output := make([]string, len(vals))

	for i, val := range vals {
		trimmedInput := strings.TrimSpace(val)
		n := len(trimmedInput)
		if n <= 3 {
			output[i] = trimmedInput
			break
		}
		output[i] = fmt.Sprintf("%c%d%c", trimmedInput[0], len(trimmedInput)-2, trimmedInput[n-1])

	}
	return output
}

func main() {
	fmt.Fprint(out, numeronyms("mango", "app"))
}
