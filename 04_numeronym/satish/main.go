package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	// make a copy of input
	cpy := make([]string, len(vals))
	for i := 0; i < len(vals); i++ {
		strLen := len(vals[i])
		if strLen < 4 {
			cpy[i] = vals[i]
		} else {
			firstChar := string(vals[i][0])
			lastChar := string(vals[i][strLen-1])
			remainingChars := strconv.Itoa(strLen - 2)
			shortForm := firstChar + remainingChars + lastChar
			cpy[i] = shortForm
		}
	}
	return cpy
}

func main() {
	vals := []string{"accessibility", "Kubernetes", "abc"}
	fmt.Fprint(out, numeronyms(vals...))
}
